#!/bin/bash

source /hab/sup/default/SystemdEnvironmentFile.sh
echo  '
[path]      
  repo = "/mnt/automate_backups/opensearch" ' > es_config.toml

hab config apply automate-ha-opensearch.default $(date '+%s') es_config.toml