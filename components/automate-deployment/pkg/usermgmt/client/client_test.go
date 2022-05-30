package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	local_users_api "github.com/chef/automate/api/interservice/local_user"
	"github.com/chef/automate/api/interservice/teams"
	"github.com/chef/automate/lib/grpc/grpctest"
	"github.com/chef/automate/lib/grpc/secureconn"
	"github.com/chef/automate/lib/tls/test/helpers"
)

func TestUserMgmtClient(t *testing.T) {
	ctx := context.Background()

	// Set up mocked out (but still functional e2e GRPC) teams-service and local-user service server
	// and related client. Mocks are generated by
	// https://github.com/chef/automate/blob/master/components/automate-grpc/protoc-gen-grpc-mock/README.md
	serviceCerts := helpers.LoadDevCerts(t, "teams-service")

	mockTeams := teams.NewTeamsServiceServerMock()
	connFactory := secureconn.NewFactory(*serviceCerts)
	g := connFactory.NewServer()
	teams.RegisterTeamsServiceServer(g, mockTeams)
	teamsServer := grpctest.NewServer(g)
	defer teamsServer.Close()

	mockLocalUser := local_users_api.NewUsersMgmtServiceServerMock()
	serviceCerts = helpers.LoadDevCerts(t, "local-user-service")
	connFactory = secureconn.NewFactory(*serviceCerts)
	g = connFactory.NewServer()
	local_users_api.RegisterUsersMgmtServiceServer(g, mockLocalUser)
	localUser := grpctest.NewServer(g)
	defer localUser.Close()

	serviceCerts = helpers.LoadDevCerts(t, "deployment-service")
	connFactory = secureconn.NewFactory(*serviceCerts)

	testClient, err := NewUserMgmtClient(ctx, connFactory, localUser.URL, teamsServer.URL)
	require.Nil(t, err)

	t.Run("CreateUser", func(t *testing.T) {
		name := "testname"
		email := "test@email.com"
		userPass := getPassword()
		expectedID := "a7cd7698-6c8a-4720-970e-5d8315caa92c" // random valid UUID

		t.Run("when the user doesn't exist it properly creates the user", func(t *testing.T) {
			// Mock CreateUser GRPC response
			mockLocalUser.GetUserFunc = func(
				context.Context, *local_users_api.Email) (*local_users_api.User, error) {
				return nil, status.Error(codes.FailedPrecondition,
					"GetUser is unexpected to be called for the code path in question.")
			}

			t.Run("and the request to create the user succeeds it returns the user's ID and true", func(t *testing.T) {
				// Mock CreateUser GRPC response
				mockLocalUser.CreateUserFunc = func(
					_ context.Context, req *local_users_api.CreateUserReq) (*local_users_api.User, error) {

					// Assert that the values passed to testClient.CreateUser
					// are the values that the local-user-service GRPC server receives.
					assert.Equal(t, name, req.Name)
					assert.Equal(t, email, req.Email)
					assert.Equal(t, userPass, req.Password)

					return &local_users_api.User{
						Name:  req.Name,
						Email: req.Email,
						Id:    expectedID,
					}, nil
				}

				responseID, wasCreated, err := testClient.CreateUser(ctx, name, email, userPass)
				require.NoError(t, err)
				assert.Equal(t, expectedID, responseID)
				assert.True(t, wasCreated)
			})

			t.Run("and the request to create the user fails it returns the original error", func(t *testing.T) {
				mockLocalUser.CreateUserFunc = func(
					_ context.Context, req *local_users_api.CreateUserReq) (*local_users_api.User, error) {

					assert.Equal(t, name, req.Name)
					assert.Equal(t, email, req.Email)
					assert.Equal(t, userPass, req.Password)

					return nil, status.Error(codes.Internal, "unexpected error")
				}

				responseID, wasCreated, err := testClient.CreateUser(ctx, name, email, userPass)
				require.Equal(t, "", responseID)
				require.Error(t, err)
				assert.False(t, wasCreated)
				grpctest.AssertCode(t, codes.Internal, err)
			})
		})

		t.Run("when the user already exists", func(t *testing.T) {
			mockLocalUser.CreateUserFunc = func(
				_ context.Context, req *local_users_api.CreateUserReq) (*local_users_api.User, error) {

				assert.Equal(t, name, req.Name)
				assert.Equal(t, email, req.Email)
				assert.Equal(t, userPass, req.Password)

				return nil, status.Error(codes.AlreadyExists, "already exists")
			}

			t.Run("and the request to get the user succeeds it returns the id, false and no error", func(t *testing.T) {
				mockLocalUser.GetUserFunc = func(
					_ context.Context, req *local_users_api.Email) (*local_users_api.User, error) {

					assert.Equal(t, email, req.Email)

					return &local_users_api.User{
						Name:  name,
						Email: req.Email,
						Id:    expectedID,
					}, nil
				}

				responseID, wasCreated, err := testClient.CreateUser(ctx, name, email, userPass)
				require.NoError(t, err)
				assert.False(t, wasCreated)
				assert.Equal(t, expectedID, responseID)
			})

			t.Run("and the request to get the user fails it returns the original error", func(t *testing.T) {
				mockLocalUser.GetUserFunc = func(
					_ context.Context, req *local_users_api.Email) (*local_users_api.User, error) {

					assert.Equal(t, email, req.Email)

					return nil, status.Error(codes.Internal, "unexpected error")
				}

				responseID, wasCreated, err := testClient.CreateUser(ctx, name, email, userPass)
				require.Equal(t, "", responseID)
				assert.False(t, wasCreated)
				require.Error(t, err)
				grpctest.AssertCode(t, codes.Internal, err)
			})
		})
	})

	t.Run("AddUserToAdminTeam", func(t *testing.T) {
		// random valid UUIDs
		teamID := "066e01e2-7d41-4b1b-ba6f-5fde7b4a6db1"
		userID := "a7cd7698-6c8a-4720-970e-5d8315caa92c"

		t.Run("when the admins team can't be retrieved it returns the original error", func(t *testing.T) {
			mockTeams.GetTeamFunc = func(
				_ context.Context, req *teams.GetTeamReq) (*teams.GetTeamResp, error) {

				assert.Equal(t, req.Id, "admins")
				return nil, status.Error(codes.Internal, "unexpected error")
			}

			err := testClient.AddUserToAdminTeam(ctx, userID)
			require.Error(t, err)
			grpctest.AssertCode(t, codes.Internal, err)
		})

		t.Run("when the admins team is properly retrieved", func(t *testing.T) {
			mockTeams.GetTeamFunc = func(
				_ context.Context, req *teams.GetTeamReq) (*teams.GetTeamResp, error) {

				assert.Equal(t, req.Id, "admins")
				return &teams.GetTeamResp{
					Team: &teams.Team{
						Id:   teamID,
						Name: "admins",
					},
				}, nil
			}

			t.Run("when add users succeeds it returns nil", func(t *testing.T) {
				mockTeams.AddTeamMembersFunc = func(
					_ context.Context, req *teams.AddTeamMembersReq) (*teams.AddTeamMembersResp, error) {

					assert.Equal(t, req.Id, teamID)
					assert.Equal(t, req.UserIds, []string{userID})

					return &teams.AddTeamMembersResp{
						UserIds: []string{userID},
					}, nil
				}

				err := testClient.AddUserToAdminTeam(ctx, userID)
				require.NoError(t, err)
			})

			t.Run("when add users fails it returns the original error", func(t *testing.T) {
				mockTeams.AddTeamMembersFunc = func(
					_ context.Context, req *teams.AddTeamMembersReq) (*teams.AddTeamMembersResp, error) {

					assert.Equal(t, req.Id, teamID)
					assert.Equal(t, req.UserIds, []string{userID})

					return nil, status.Error(codes.Internal, "unexpected error")
				}

				err := testClient.AddUserToAdminTeam(ctx, userID)
				require.Error(t, err)
				grpctest.AssertCode(t, codes.Internal, err)
			})
		})
	})
}

func getPassword() string {
	return "GottaCatchEmAll"
}
