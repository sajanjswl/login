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


variable "ecr_image_api" {
  description = "ECR Image for API"
  default     = "public.ecr.aws/f0x8s9w9/aws-user-service:latest"
}

variable "ecr_image_proxy" {
  description = "ECR Image for API"
  default     = "public.ecr.aws/f0x8s9w9/aws-user-service:proxy"
}


