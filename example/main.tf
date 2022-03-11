provider "proxmox" {
  virtual_environment {
    endpoint = var.virtual_environment_endpoint
    username = var.virtual_environment_username
    password = var.virtual_environment_password
    tokenname = var.virtual_environment_tokenname
    tokenvalue = var.virtual_environment_tokenvalue
    insecure = true
  }
}
