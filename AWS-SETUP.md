

##AWS SETUP
Use this  region (US East (N. Virginia)us-east-1) 
* Login to you aws account using `Adminstrator IAM User` and  create a private s3 bucket (enable bucket versioning)
for storing terrafom state.(In use `user-service-aws-terraform-state-bucket`)

* Create a DynamoDB table (in use`user-service-aws-terraform-state-lock
` ) PrimaryKey=`LockID`  rest default. This table is to  put a lock on infastructure while it's in use. So that only one istance of terraform can be applied at any point in time.
  
* Create an AWS-ECR for storing images enable scan on push (in use `user-service-aws-terraform-registry` ) 

* AWS bastion instance ssh key setup: 
* ec2 dashboard -> network and security -> Key Pairs
* provide the key  pair name to `bastion_key_name` in variable.tf file

# Regenrate vault-credentials
aws-vault exec sajan.jaiswal --duration=12h


# Terraform workspace
## Bastion Instance
An isnstance profile is something that we can assign to our bastion isnstance to give it IAM role information

We are creating a bastion  instance with the ability to assume a role and then we will be creating IAM role that will allow it have access to our ecr repository


Security Group allow you to to control the inbound and outbound access  allowed to that resource


# ssh into bastion instance
ssh ec2-user@ec2-54-224-151-134.compute-1.amazonaws.com


#postgres commands
docker run -it --rm \
    --network user-network \
    bitnami/postgresql:latest psql -h postgres -U postgres

psql -h $(terraform output -raw rds_hostname) -p $(terraform output -raw rds_port) -U $(terraform output -raw rds_username) postgres < /Users/sajanjswl/sandbox/dl-auth/backup.sql


psql -h education.clzpfr8xlxjo.us-east-2.rds.amazonaws.com -p 5432 -U edu postgres


docker run -it --rm jbergknoff/postgresql-client postgresql://recihgjhghpeapp:changeme67r6fvfy@user-service-dev-db.c7bjzqhwgcal.us-east-1.rds.amazonaws.com:5432/user-service


db_username = "recihgjhghpeapp"
db_password = "changeme67r6fvfy"
db_name     = "user-service"



 # on bastion
docker run -it --rm jbergknoff/postgresql-client postgresql://recihgjhghpeapp:changeme67r6fvfy@user-service-default-db.c7bjzqhwgcal.us-east-1.rds.amazonaws.com:5432/userservice