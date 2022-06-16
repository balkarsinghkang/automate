+++
title = "Prerequisites"

draft = false

gh_repo = "automate"

[menu]
  [menu.automate]
    title = "Prerequisites"
    identifier = "automate/deploy_high_availability/backup_and_restore/ha_backup_restore_prerequisites.md Prerequisites"
    parent = "automate/deploy_high_availability/backup_and_restore"
    weight = 210
+++

This page explains the prerequisites of the AWS and on-premise backup.

{{< note >}} You can take backup on EFS system through DNS or IP. {{< /note >}}

## AWS Backed Backup

The two pre-backup configurations for AWS are:

- For S3 Backup
- For EFS Backup

### Pre Backup Configuration for S3 Backup

To run the terraform scripts, the IAM users should have proper permissions. Going forward, we will also discuss the required permissions. You should have your  `secret access key` and `key id`, or else you can regenerate a new access key.

**Permissions Required**

Check if the IAM user has all the required permissions. The permission policies are listed below:

1. AdministratorAccess

1. APIGatewayAdministrator (For aws AmazonAPIGatewayAdministrator)

1. S3FullAccess (for aws AmazonS3FullAccess)

Create an IAM role to give access of **S3** to **OpenSearch** instances. The role should already be assigned as the OpenSearch instance tries to communicate S3.

The permissions can either be directly added to the user or added via **IAM Group**.

Once done with the above steps, `.toml` file and patch the `.config`. In the file, modify the values listed below:

1. **bucket name (bucket = "bucket-name" and name = "bucket-name")**

2. `mkdir configs`

3. `vi configs/automate.toml`

Refer to the content for the `automate.toml` file below:

```sh
[global.v1.external.opensearch.backup]
    enable = true
    location = "s3"

[global.v1.external.opensearch.backup.s3]

    # bucket (required): The name of the bucket
    bucket = "bucket-name"

    # base_path (optional):  The path within the bucket where backups should be stored
    # If base_path is not set, backups will be stored at the root of the bucket.
    base_path = "opensearch"

    # name of an s3 client configuration you create in your opensearch.yml
    # see https://www.open.co/guide/en/opensearch/plugins/current/repository-s3-client.html
    # for full documentation on how to configure client settings on your
    # OpenSearch nodes
    client = "default"

[global.v1.external.opensearch.backup.s3.settings]
    ## The meaning of these settings is documented in the S3 Repository Plugin
    ## documentation. See the following links:
    ## https://www.open.co/guide/en/opensearch/plugins/current/repository-s3-repository.html

    ## Backup repo settings
    # compress = false
    # server_side_encryption = false
    # buffer_size = "100mb"
    # canned_acl = "private"
    # storage_class = "standard"
    ## Snapshot settings
    # max_snapshot_bytes_per_sec = "40mb"
    # max_restore_bytes_per_sec = "40mb"
    # chunk_size = "null"
    ## S3 client settings
    # read_timeout = "50s"
    # max_retries = 3
    # use_throttle_retries = true
    # protocol = "https"

[global.v1.backups]
    location = "s3"

[global.v1.backups.s3.bucket]
    # name (required): The name of the bucket
    name = "bucket-name"

    # endpoint (required): The endpoint for the region the bucket lives in.
    # See https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
    endpoint = "https://s3.amazonaws.com"

    # base_path (optional):  The path within the bucket where backups should be stored
    # If base_path is not set, backups will be stored at the root of the bucket.
    base_path = "automate"

[global.v1.backups.s3.credentials]
    access_key = "AKIARUQHMSKV6BXLAXHO"
    secret_key = "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"
```

Execute the command given below to trigger the deployment.

```sh
./chef-automate config patch configs/automate.toml
```

Take a [back-up](/automate/ha_backup_for_backup_restore) of the configurations once the cluster has been deployed.

{{< note >}} **IAM Role:** Assign the IAM Role to all the OpenSearch instances in the cluster created above. {{< /note >}}

### Elastic File System(EFS) Configuration for backup

To backup on the shared file system for AWS, follow the steps given below:

- Create the **EFS** over AWS.
- Mount the creating EFS using **DNS** or **IP**.
- Open the **Port(2049) Proto(NFS)** for EFS Security Group.

## On-Premise Backed Backup

The two pre-backup configurations for On-Premise are:

- For File System Backup
- For Object Storage

### Pre Backup Configuration for File System Backup

A shared file system is always required to create **OpenSearch** snapshots. To register the snapshot repository using OpenSearch, it is necessary to mount the same shared filesystem to the exact location on all master and data nodes. Register the location (or one of its parent directories) in the `path.repo` setting on all master and data nodes.

Once the shared filesystem is mounted to `/mnt/automate_backups`, configure Automate to register the snapshot locations with OpenSearch.

- Mount the shared file system on all OpenSearch servers:

```sh
mount /mnt/automate_backups
```

- Create an OpenSearch sub-directory and set permissions to one of the OpenSearch server (only if the network mount is correctly mounted).

```sh
sudo mkdir /mnt/automate_backups/opensearch
sudo chown hab:hab /mnt/automate_backups/opensearch/
```

Configure the OpenSearch `path.repo` setting by SSH to a single OpenSearch server by following the steps given below:

- Export the current OpenSearch config from the Habitat supervisor. Get the root access to run the following commands:

```sh
source /hab/sup/default/SystemdEnvironmentFile.sh
automate-backend-ctl applied --svc=automate-ha-opensearch | tail -n +2 > es_config.toml
```

- Edit `es_config.toml` and add the following settings to the end of the file.

{{< note >}} If the credentials have never been rotated, the above file may be empty. {{< /note >}}

```sh
[es_yaml.path]   
# Replace /mnt/automate_backups with the backup_mount config found on the provisioning host in /hab/a2_deploy_workspace/a2ha.rb   
repo = "/mnt/automate_backups/opensearch" 
```

To trigger the restart of the OpenSearch on each server, apply the updated  `es_config.toml` config to OpenSearch once.

```sh
hab config apply automate-ha-opensearch.default $(date '+%s') es\_config.toml

hab svc status (check whether OpenSearch service is up or not)

curl -k -X GET "<https://localhost:9200/_cat/indices/*?v=true&s=index&pretty>" -u admin:admin (Another way to check is to check whether all the indices are green or not)

# Watch for a message about OpenSearch going from RED to GREEN
`journalctl -u hab-sup -f | grep 'automate-ha-opensearch'
```

- Configure Automate to handle *External OpenSearch Backups*.

- Create an `automate.toml` file on the provisioning server using the following command:

```bash
touch automate.toml
```

Add the following configuration to `automate.toml` on the provisioning host:

```sh
[global.v1.external.opensearch.backup]
enable = true
location = "fs"

[global.v1.external.opensearch.backup.fs]
# The `path.repo` setting you've configured on your OpenSearch nodes must be a parent directory of the setting you configure here:
path = "/mnt/automate_backups/opensearch"

[global.v1.backups.filesystem]
path = "/mnt/automate_backups/backups"
```

- Patch the `.config` to trigger the deployment.

```sh
./chef-automate config patch automate.toml
```

### Pre-Backup Configuration for Object Storage

This section provides the pre-backup configuration required to backup the data on Object Storage System (Other than AWS S3) like *Minio*, *Non-AWS S3*. The steps to set a secret key using commands are given below:

1. Log in to all the opensearch nodes and follow the steps on all the opensearch nodes.
  
  - Export `ES_PATH_CONF="/hab/svc/automate-ha-opensearch/config"`
  - Hab pkg `exec chef/opensearch-odfe opensearch-keystore add s3.client.default.access_key` (When asked, Enter your key)
  - Hab pkg `exec chef/opensearch-odfe opensearch-keystore add s3.client.default.secret_key` (When asked, Enter your key/secret)
  - Chown `hab:hab /hab/svc/automate-ha-opensearch/config/opensearch.keystore` (Setting hab:hab permission)
  - Curl `-k -X POST "https://127.0.0.1:9200/_nodes/reload_secure_settings?pretty" -u admin:admin` (Command to load the above setting)

The final output after running the command 1.5 on the third node is given below:

```json
{
  "_nodes" : {
    "total" : 3,
    "successful" : 3,
    "failed" : 0
  },
  "cluster_name" : "chef-insights",
  "nodes" : {
    "lenRTrZ1QS2uv_vJIwL-kQ" : {
      "name" : "lenRTrZ"
    },
    "Us5iBo4_RoaeojySjWpr9A" : {
      "name" : "Us5iBo4"
    },
    "qtz7KseqSlGm2lEm0BiUEg" : {
      "name" : "qtz7Kse"
    }
  }
}
```

2. To override the existing default endpoint:

- Login to one of the open search instances and run the following command (You need root access to run the command):

```sh
source /hab/sup/default/SystemdEnvironmentFile.sh
automate-backend-ctl applied --svc=automate-ha-opensearch | tail -n +2 > es_config.toml
```

- Edit the created `es_config.toml` file and add the following settings at the end of the file. (*The file will be empty if the credentials have not been rotated*)

```sh
[es_yaml.s3.client.default]
endpoint = "< S3 endpoint, e.g. xyz.s3.com>"
```

- Run the following command to apply the updated `es_config.toml` changes. Run this command only once. (*This will trigger a restart of the OpenSearch services on each server*)

```sh
hab config apply automate-ha-opensearch.default $(date '+%s') es_config.toml
```

- Once done with the above steps, run the following command:

```sh
journalctl -u hab-sup -f | grep 'automate-ha-opensearch'
```

The screen will display a message of OpenSearch going from **RED/YELLOW** to **GREEN**.