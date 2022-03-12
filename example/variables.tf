variable "virtual_environment_endpoint" {
  type        = string
  description = "The endpoint for the Proxmox Virtual Environment API (example: https://host:port)"
}

variable "virtual_environment_password" {
  type        = string
  description = "The password for the Proxmox Virtual Environment API"
}

variable "virtual_environment_username" {
  type        = string
  description = "The username and realm for the Proxmox Virtual Environment API (example: root@pam)"
}

variable "virtual_environment_sshkey" {
  type        = string
  description = "The filename of the SSH key to authorize the connection to Proxmox via SSH (example: /home/me/.ssh/id_rsa)"
}

variable "virtual_environment_tokenname" {
  type        = string
  description = "The optional name of the API token for the Proxmox Virtual Environment API (example: provisioning)"
  default     = ""
}

variable "virtual_environment_tokenvalue" {
  type        = string
  description = "The optional value of the API token for the Proxmox Virtual Environment API (example: 12345678-1234-1234-1234-1234567890ab)"
  default     = ""
}
