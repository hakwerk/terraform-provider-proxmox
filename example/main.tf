provider "proxmox" {
  virtual_environment {
    endpoint = var.virtual_environment_endpoint
    username = var.virtual_environment_username
    #password = var.virtual_environment_password
    sshkey = var.virtual_environment_sshkey
    tokenname = var.virtual_environment_tokenname
    tokenvalue = var.virtual_environment_tokenvalue
    insecure = true
  }
}
