variable "airgap_info" {
}

variable "backend_aib_dest_file" {
}

variable "backend_aib_local_file" {
}

variable "backup_config_efs" {
  default = "false"
}

variable "curator_pkg_ident" {
}

variable "elasticsearch_pkg_ident" {
}

variable "elasticsearch_svc_load_args" {
}

variable "elasticsidecar_pkg_ident" {
}

variable "elasticsidecar_svc_load_args" {
}

variable "habitat_info" {
}

variable "journalbeat_pkg_ident" {
}

variable "kibana_pkg_ident" {
}

variable "metricbeat_pkg_ident" {
}

variable "opensearch_instance_count" {
  default = 3
}

variable "opensearch_listen_port" {
  default = 9200
}

variable "private_ips" {
  default = []
}

variable "public_ips" {
  default = []
}

variable "ssh_key_file" {
}

variable "ssh_user" {
  default = "centos"
}

variable "ssh_user_sudo_password" {
}

variable "sudo_cmd" {
  default = "sudo"
}

variable "tmp_path" {
  default = "/var/automate-ha"
}
