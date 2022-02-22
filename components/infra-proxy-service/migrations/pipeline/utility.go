package pipeline

import (
	"archive/zip"
	"context"
	"encoding/json"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/chef/automate/api/interservice/authz"
	"github.com/chef/automate/components/infra-proxy-service/pipeline"
	"github.com/chef/automate/components/infra-proxy-service/storage"
	log "github.com/sirupsen/logrus"
)

// StoreOrgs reads the Result struct and populate the orgs table
func StoreOrgs(ctx context.Context, st storage.Storage, mst storage.MigrationStorage, authzProjectClient authz.ProjectsServiceClient, res pipeline.Result) (pipeline.Result, error) {
	var err error
	var msg string
	var totalSucceeded, totalSkipped, totalFailed int64
	_, err = mst.StartOrgMigration(ctx, res.Meta.MigrationID, res.Meta.ServerID)
	if err != nil {
		return res, err
	}
	log.Info("Starting the organisation migration phase for migration id: ", res.Meta.MigrationID)
	for _, org := range res.ParsedResult.Orgs {
		err, _ = StoreOrg(ctx, st, org, res.Meta.ServerID, authzProjectClient)
		if err != nil {
			totalFailed++
			msg = err.Error()
			continue
		}
		if org.ActionOps == pipeline.Skip {
			totalSkipped++
			continue
		}
		totalSucceeded++
	}
	if len(res.ParsedResult.Orgs) == int(totalFailed) {
		log.Errorf("Failed to migrate orgs for migration id %s : %s", res.Meta.MigrationID, err.Error())
		_, _ = mst.FailedOrgMigration(ctx, res.Meta.MigrationID, res.Meta.ServerID, msg, totalSucceeded, totalSkipped, totalFailed)
		return res, err
	}
	_, err = mst.CompleteOrgMigration(ctx, res.Meta.MigrationID, res.Meta.ServerID, totalSucceeded, totalSkipped, totalFailed)
	if err != nil {
		log.Errorf("Failed to update the status for migration id %s : %s", res.Meta.MigrationID, err.Error())
		return res, err
	}
	log.Info("Successfully completed the organisation migration phase for migration id: ", res.Meta.MigrationID)
	return res, err
}

// StoreOrg stores a single Org into DB
func StoreOrg(ctx context.Context, st storage.Storage, org pipeline.Org, serverID string, authzProjectClient authz.ProjectsServiceClient) (error, pipeline.ActionOps) {
	var actionTaken pipeline.ActionOps
	var err error
	switch org.ActionOps {
	case pipeline.Insert:
		projects, err := createProjectFromOrgIdAndServerID(ctx, serverID, org.Name, authzProjectClient)
		if err != nil {
			log.Errorf("Unable to create project for serverid: %s", serverID)
			return err, actionTaken
		}
		_, err = st.StoreOrg(ctx, org.Name, org.FullName, "", "", serverID, projects)
		if err != nil {
			log.Errorf("Unable to insert org for server id: %s", serverID)
		}
		actionTaken = pipeline.Insert
	case pipeline.Delete:
		_, err = st.DeleteOrg(ctx, org.Name, serverID)
		actionTaken = pipeline.Delete
	case pipeline.Update:
		_, err = st.EditOrg(ctx, org.Name, org.FullName, "", serverID, nil)
		actionTaken = pipeline.Update
	default:
	}
	return err, actionTaken
}

//function to create a new iam project for each client
func createProjectFromOrgIdAndServerID(ctx context.Context, serverId string, orgId string, authzProjectClient authz.ProjectsServiceClient) ([]string, error) {

	newProject := &authz.CreateProjectReq{
		Name:         serverId + "_" + orgId,
		Id:           serverId + "_" + orgId,
		SkipPolicies: false,
	}

	projectID, err := authzProjectClient.CreateProject(ctx, newProject)
	if err != nil {
		return nil, err
	}

	return []string{projectID.Project.Name}, nil
}

func ParseOrgs(ctx context.Context, st storage.Storage, mst storage.MigrationStorage, result pipeline.Result) (pipeline.Result, error) {
	var err error
	log.Info("Starting with organisation parsing phase for migration id: ", result.Meta.MigrationID)
	_, err = mst.StartOrgParsing(ctx, result.Meta.MigrationID, result.Meta.ServerID)
	if err != nil {
		log.Errorf("Failed to update the status for start org parssing for the migration id path %s : %s", result.Meta.MigrationID, err.Error())
		return result, err
	}
	orgPath := path.Join(result.Meta.UnzipFolder, "organizations")
	folder, err := os.Open(orgPath)
	if err != nil {
		log.Errorf("Failed to open the folder for the file path %s : %s", orgPath, err.Error())
		_, _ = mst.FailedOrgParsing(ctx, result.Meta.MigrationID, result.Meta.ServerID, err.Error(), 0, 0, 0)
		return result, err
	}

	orgNames, err := folder.Readdir(0)
	if err != nil {
		log.Errorf("Failed to read the files for the file path %s : %s", orgPath, err.Error())
		_, _ = mst.FailedOrgParsing(ctx, result.Meta.MigrationID, result.Meta.ServerID, err.Error(), 0, 0, 0)
		return result, err
	}
	_ = folder.Close()
	orgsPresentInDB, err := st.GetOrgs(ctx, result.Meta.ServerID)
	if err != nil {
		log.Errorf("Failed to read orgs from database for %s:%s", result.Meta.ServerID, err.Error())
		_, _ = mst.FailedOrgParsing(ctx, result.Meta.MigrationID, result.Meta.ServerID, err.Error(), 0, 0, 0)
		return result, err
	}

	result.ParsedResult.Orgs = append(result.ParsedResult.Orgs, insertOrUpdateOrg(orgNames, orgsPresentInDB, orgPath)...)

	result.ParsedResult.Orgs = append(result.ParsedResult.Orgs, deleteOrgsIfNotPresentInCurrentFile(orgNames, orgsPresentInDB)...)

	_, err = mst.CompleteOrgParsing(ctx, result.Meta.MigrationID, result.Meta.ServerID, 0, 0, 0)
	if err != nil {
		log.Errorf("Failed to update the complete status while parsing for migration id %s : %s", result.Meta.MigrationID, err.Error())
		return result, err
	}

	log.Info("Successfully completed the organisation parsing phase for migration id: ", result.Meta.MigrationID)
	return result, nil

}

//CreatePreview Stores the staged data in db
func CreatePreview(ctx context.Context, st storage.Storage, mst storage.MigrationStorage, result pipeline.Result) (pipeline.Result, error) {
	log.Info("Starting with create preview phase for migration id: ", result.Meta.MigrationID)

	_, err := mst.StartCreatePreview(ctx, result.Meta.MigrationID, result.Meta.ServerID)
	if err != nil {
		log.Errorf("Failed to update the status for create preview for the migration id  %s : %s", result.Meta.MigrationID, err.Error())
		return result, err
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		log.Errorf("Failed to marshal the staged data for the migration id %s : %s", result.Meta.MigrationID, err.Error())
		_, _ = mst.FailedCreatePreview(ctx, result.Meta.MigrationID, result.Meta.ServerID, err.Error(), 0, 0, 0)
		return result, err
	}

	_, err = mst.StoreMigrationStage(ctx, result.Meta.MigrationID, resultByte)
	if err != nil {
		log.Errorf("Failed to store the staged data %s : %s ", result.Meta.MigrationID, err.Error())
		_, _ = mst.FailedCreatePreview(ctx, result.Meta.MigrationID, result.Meta.ServerID, err.Error(), 0, 0, 0)
		return result, err
	}
	_, err = mst.CompleteCreatePreview(ctx, result.Meta.MigrationID, result.Meta.ServerID, 0, 0, 0)
	if err != nil {
		log.Errorf("Failed to update the complete status while creating preview for migration id %s : %s", result.Meta.MigrationID, err.Error())
		return result, err
	}
	log.Info("Successfully completed the create preview phase for migration id: ", result.Meta.MigrationID)

	return result, nil
}

func createDatabaseOrgsMap(orgs []storage.Org) map[string]string {
	orgMap := make(map[string]string)
	for _, s := range orgs {
		orgMap[s.ID] = s.Name
		//Values for comparison
	}
	return orgMap
}

func createFileOrgsMap(orgs []os.FileInfo) map[string]string {
	orgMap := make(map[string]string)
	for _, s := range orgs {
		if s.IsDir() {
			orgMap[s.Name()] = ""
			//No value required for comparison
		}
	}
	return orgMap
}

func insertOrUpdateOrg(orgsInFiles []os.FileInfo, orgsInDB []storage.Org, orgPath string) []pipeline.Org {
	var orgList []pipeline.Org
	orgDatabaseMap := createDatabaseOrgsMap(orgsInDB)
	var orgJson pipeline.OrgJson
	log.Info("Comparing the organisations from database and backup file for insert,update and skip action")
	//For insert, update and skip action
	for _, org := range orgsInFiles {
		if org.IsDir() {
			orgInfo, valuePresent := orgDatabaseMap[org.Name()]
			orgJson = openOrgFolder(org, orgPath)
			if valuePresent {
				if orgJson.FullName != orgInfo {
					//Update org in the result actions
					orgList = append(orgList, createOrgStructForAction(orgJson.Name, orgJson.FullName, pipeline.Update))
				} else {
					//Skip org action if full names are not equal
					orgList = append(orgList, createOrgStructForAction(orgJson.Name, orgJson.FullName, pipeline.Skip))

				}
			} else {
				//Insert org action if not present in database
				orgList = append(orgList, createOrgStructForAction(orgJson.Name, orgJson.FullName, pipeline.Insert))
			}
		}
	}
	log.Info("Completed comparing the organisations from database and backup file for insert,update and skip action")
	return orgList
}

func deleteOrgsIfNotPresentInCurrentFile(orgsInFiles []os.FileInfo, orgsInDB []storage.Org) []pipeline.Org {
	var orgList []pipeline.Org
	orgFilesMap := createFileOrgsMap(orgsInFiles)
	log.Info("Comparing the organisations from database and backup file for delete action")
	//For delete action by comparing database orgs with file orgs
	for _, org := range orgsInDB {
		_, valuePresent := orgFilesMap[org.ID]
		if !valuePresent {
			orgList = append(orgList, createOrgStructForAction(org.ID, org.Name, pipeline.Delete))
		}
	}
	log.Info("Completed comparing the organisations from database and backup file for delete action")
	return orgList
}

func openOrgFolder(org os.FileInfo, fileLocation string) pipeline.OrgJson {
	var orgJson pipeline.OrgJson
	jsonPath := path.Join(fileLocation, org.Name(), "org.json")
	jsonFile, err := os.Open(jsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Errorf("Unable to open the file at location : %s", jsonPath)
	}
	log.Info("Successfully opened the file at location", jsonPath)
	defer func() {
		_ = jsonFile.Close()
	}()
	// defer the closing of our jsonFile so that we can parse it later on
	_ = json.NewDecoder(jsonFile).Decode(&orgJson)
	return orgJson
}

func createOrgStructForAction(orgId string, orgName string, ops pipeline.ActionOps) pipeline.Org {
	return pipeline.Org{
		Name:      orgId,
		FullName:  orgName,
		ActionOps: ops,
	}
}

// Unzip will decompress a zip file and sets the UnzipFolder
func Unzip(ctx context.Context, mst storage.MigrationStorage, result pipeline.Result) (pipeline.Result, error) {

	var fpath string
	_, err := mst.StartUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID)
	if err != nil {
		log.Errorf("Failed to update status in DB: %s :%s", result.Meta.MigrationID, err)
	}

	reader, err := zip.OpenReader(result.Meta.ZipFile)
	if err != nil {
		log.Errorf("cannot open reader: %s.", err)
		_, err := mst.FailedUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, "cannot open zipfile", 0, 0, 0)
		if err != nil {
			log.Errorf("Failed to update status in DB: %s", err)
		}
		return result, err
	}

	for _, file := range reader.File {

		fpath = filepath.Join(filepath.Dir(result.Meta.ZipFile), file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			log.Errorf("cannot create directory: %s. ", err)
			mst.FailedUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, "cannot create directory", 0, 0, 0)
			return result, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			file.Mode())
		if err != nil {
			log.Errorf("cannot create a file: %s.", err)
			mst.FailedUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, "cannot create a file", 0, 0, 0)
			return result, err
		}

		readClose, err := file.Open()
		if err != nil {
			log.Errorf("cannot open file")
			mst.FailedUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, "cannot open file", 0, 0, 0)
			return result, err
		}

		_, err = io.Copy(outFile, readClose)
		if err != nil {
			log.Errorf("cannot copy file")
			mst.FailedUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, "cannot copy file", 0, 0, 0)
			return result, err
		}

		outFile.Close()
		readClose.Close()

		if err != nil {
			log.Errorf("cannot copy a file")
			mst.FailedUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, "cannot copy a file", 0, 0, 0)
			return result, err
		}
	}

	result.Meta.UnzipFolder = filepath.Dir(fpath)
	reader.Close()
	_, err = mst.CompleteUnzip(ctx, result.Meta.MigrationID, result.Meta.ServerID, 0, 0, 0)
	if err != nil {
		log.Errorf("Failed to update status in DB: %s :%s", result.Meta.MigrationID, err)
	}
	return result, nil
}

// ParseOrgUserAssociation  sync the automate org users with chef server org users
func ParseOrgUserAssociation(ctx context.Context, st storage.Storage, result pipeline.Result) (pipeline.Result, error) {
	log.Info("Starting with the parsing org user association for migration id :", result.Meta.MigrationID)
	var orgUserAssociations []pipeline.OrgsUsersAssociations
	var err error
	orgUserAssociations, err = getActionForOrgUsers(ctx, st, result)
	if err != nil {
		log.Errorf("Unable to parse org user association for migration id : %s : %s", result.Meta.MigrationID, err.Error())
		return result, err
	}
	result.ParsedResult.OrgsUsers = append(result.ParsedResult.OrgsUsers, orgUserAssociations...)
	log.Info("Completed with the parsing org user association for migration id :", result.Meta.MigrationID)
	return result, nil
}

func getActionForOrgUsers(ctx context.Context, st storage.Storage, result pipeline.Result) ([]pipeline.OrgsUsersAssociations, error) {
	orgUserAssociations := make([]pipeline.OrgsUsersAssociations, 0)
	var userAssociations []pipeline.UserAssociation
	orgPath := path.Join(result.Meta.UnzipFolder, "organizations")
	for _, org := range result.ParsedResult.Orgs {
		log.Info("Getting actions for org id", org.Name)
		chefServerOrgUsers, err := getChefServerOrgUsers(org.Name, orgPath)
		if err != nil {
			log.Errorf("Unable to get the chef server organisation users %s ", err)
			return nil, err
		}
		if org.ActionOps == pipeline.Insert {
			userAssociations = append(userAssociations, createInsertUserAssociation(chefServerOrgUsers)...)
		} else {
			if org.ActionOps == pipeline.Delete {
				userAssociations = append(userAssociations, createDeleteUserAssociation(chefServerOrgUsers)...)
			} else {
				orgUsersInDb, err := st.GetAutomateOrgUsers(ctx, org.Name)
				if err != nil {
					log.Errorf("Unable to fetch automate Users for org %s : %s", org.Name, err.Error())
					return nil, err
				}
				userAssociations = append(userAssociations, insertOrUpdateActionForOrgUsers(orgUsersInDb, chefServerOrgUsers)...)
				userAssociations = append(userAssociations, deleteActionForOrgUses(orgUsersInDb, chefServerOrgUsers)...)
			}
		}
		orgUserAssociations = append(orgUserAssociations, pipeline.OrgsUsersAssociations{OrgName: org, Users: userAssociations})
	}
	return orgUserAssociations, nil
}

func createInsertUserAssociation(chefServerOrgUsers []pipeline.UserAssociation) []pipeline.UserAssociation {
	userAssociation := make([]pipeline.UserAssociation, 0)
	for _, user := range chefServerOrgUsers {
		userAssociation = append(userAssociation, pipeline.UserAssociation{Username: user.Username, IsAdmin: user.IsAdmin, ActionOps: pipeline.Insert})
	}
	return userAssociation
}

func createDeleteUserAssociation(chefServerOrgUsers []pipeline.UserAssociation) []pipeline.UserAssociation {
	userAssociation := make([]pipeline.UserAssociation, 0)
	for _, user := range chefServerOrgUsers {
		userAssociation = append(userAssociation, pipeline.UserAssociation{Username: user.Username, IsAdmin: user.IsAdmin, ActionOps: pipeline.Delete})
	}
	return userAssociation
}

func insertOrUpdateActionForOrgUsers(orgUsers []storage.OrgUser, chefServerOrgUsers []pipeline.UserAssociation) []pipeline.UserAssociation {
	var userAssociation []pipeline.UserAssociation
	orgUserMapDB := createMapForOrgUsersInDB(orgUsers)
	for _, user := range chefServerOrgUsers {
		isAdmin, valuePresent := orgUserMapDB[user.Username]
		if valuePresent {
			//check for the org admins
			if user.IsAdmin != isAdmin {
				userAssociation = append(userAssociation, pipeline.UserAssociation{Username: user.Username, IsAdmin: user.IsAdmin, ActionOps: pipeline.Update})
			} else {
				userAssociation = append(userAssociation, pipeline.UserAssociation{Username: user.Username, IsAdmin: user.IsAdmin, ActionOps: pipeline.Skip})
			}
		} else {
			userAssociation = append(userAssociation, pipeline.UserAssociation{Username: user.Username, IsAdmin: user.IsAdmin, ActionOps: pipeline.Insert})
		}
	}
	return userAssociation
}

func deleteActionForOrgUses(orgUsers []storage.OrgUser, chefServerOrgUsers []pipeline.UserAssociation) []pipeline.UserAssociation {
	var userAssociation []pipeline.UserAssociation
	orgUserJsonMap := createMapForOrgUsersInJson(chefServerOrgUsers)
	for _, user := range orgUsers {
		_, valuePresent := orgUserJsonMap[user.InfraServerUsername]
		if !valuePresent {
			userAssociation = append(userAssociation, pipeline.UserAssociation{Username: user.InfraServerUsername, IsAdmin: user.IsAdmin, ActionOps: pipeline.Delete})
		}
	}

	return userAssociation
}

func createMapForOrgUsersInDB(orgUsers []storage.OrgUser) map[string]bool {
	orgUsersMap := make(map[string]bool)
	for _, s := range orgUsers {
		orgUsersMap[s.InfraServerUsername] = s.IsAdmin
	}
	return orgUsersMap
}

func createMapForOrgUsersInJson(chefServerOrgUsers []pipeline.UserAssociation) map[string]string {
	orgUsersMap := make(map[string]string)
	for _, user := range chefServerOrgUsers {
		orgUsersMap[user.Username] = ""
	}
	return orgUsersMap
}

// getChefServerOrgUsers returns the chef server organisation users from backup file
func getChefServerOrgUsers(orgName, fileLocation string) ([]pipeline.UserAssociation, error) {
	orgUsers := make([]pipeline.UserAssociation, 0)

	members, err := getOrgMembers(orgName, fileLocation)
	if err != nil {
		log.Errorf("Unable to get orgnisation members %s", err)
		return nil, err
	}
	admins, err := getOrgAdmins(orgName, fileLocation)
	if err != nil {
		log.Errorf("Unable to get orgnisation admins %s", err)
		return nil, err
	}
	orgAdminMap := createMapForOrgAdminsInJson(admins)
	for _, member := range members {
		orgUser := pipeline.UserAssociation{}
		if member.User.Username == "pivotal" {
			continue
		}
		orgUser.Username = member.User.Username
		_, valuePresent := orgAdminMap[member.User.Username]
		if valuePresent {
			orgUser.IsAdmin = true
		}
		orgUsers = append(orgUsers, orgUser)
	}
	return orgUsers, nil
}

// getOrgMembers Get the data of members.json
func getOrgMembers(orgName, fileLocation string) ([]pipeline.MembersJson, error) {
	var orgMembers []pipeline.MembersJson
	usersJsonPath := path.Join(fileLocation, orgName, "members.json")
	usersjsonFile, err := os.Open(usersJsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Errorf("Unable to open org members file at the location : %s", usersJsonPath)
		return nil, err
	}
	log.Info("Successfully opened the org members file at location", usersJsonPath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		_ = usersjsonFile.Close()
	}()
	err = json.NewDecoder(usersjsonFile).Decode(&orgMembers)
	if err != nil {
		log.Errorf("Unable to decode the org members file %s %s", usersJsonPath, err)
		return nil, err
	}
	return orgMembers, nil
}

// getOrgAdmins Get the data of admins.json
func getOrgAdmins(orgName, fileLocation string) (pipeline.AdminsJson, error) {
	var orgAdmins pipeline.AdminsJson
	adminJsonPath := path.Join(fileLocation, orgName, "groups", "admins.json")
	jsonFile, err := os.Open(adminJsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Errorf("Unable to open org admins file at the location : %s %s", adminJsonPath, err)
		return pipeline.AdminsJson{}, err
	}
	log.Info("Successfully opened the org admins file at location", adminJsonPath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		_ = jsonFile.Close()
	}()
	err = json.NewDecoder(jsonFile).Decode(&orgAdmins)
	if err != nil {
		log.Errorf("Unable to decode the org admins file %s %s", adminJsonPath, err)
		return pipeline.AdminsJson{}, err
	}
	return orgAdmins, nil
}

func createMapForOrgAdminsInJson(adminsJson pipeline.AdminsJson) map[string]string {
	orgAdminsMap := make(map[string]string)
	for _, username := range adminsJson.Users {
		orgAdminsMap[username] = ""
	}
	return orgAdminsMap
}