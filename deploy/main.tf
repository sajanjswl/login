terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.37.0"
    }
  }
  backend "s3" {
    bucket         = "user-service-aws-terraform-state-bucket"
    key            = "user-service-aws-terraform-state.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "user-service-aws-terraform-state-lock"
  }
}

# terraform {
#   required_providers  {
#     aws = {
#       source = "hashicorp/aws"
#       version = "3.37.0"
#     }
#   }
# }

provider "aws" {
  region = "us-east-1"
  # version = "~> 2.54.0"
  # version = "~> 3.27.0"
  # version ="~> 3.37"
}


locals {
  prefix = "${var.prefix}-${terraform.workspace}"

  common_tags = {
    Environment = terraform.workspace
    Project     = var.project
    Owner       = var.contact
    ManagedBy   = "Terraform"
  }

}

data "aws_region" "current" {}