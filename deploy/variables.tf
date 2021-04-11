variable "prefix" {
  default = "raad"
}


variable "project" {
  default = "recipe-app-api-devops"
}

variable "contact" {
  default = "sjnjaiswal@gmail.com"
}

variable "db_username" {
  description = "Username for the RDS Postgres instance"
}

variable "db_password" {
  description = "Password for the RDS postgres instance"
}


variable "bastion_key_name" {
  default = "recipe-app-api-devops-bastion"
}

