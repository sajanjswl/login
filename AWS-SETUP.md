

##AWS SETUP
Use this  region (US East (N. Virginia)us-east-1) 
* Login to you aws account using `Adminstrator IAM User` and  create a private s3 bucket (enable bucket versioning)
for storing terrafom state.(In use `user-service-aws-terraform-state-bucket`)

* Create a DynamoDB table (in use`user-service-aws-terraform-state-lock
` ) PrimaryKey=`LockID`  rest default. This table is to  put a lock on infastructure while it's in use. So that only one istance of terraform can be applied at any point in time.
  
* Create an AWS-ECR for storing images enable scan on push (in use `user-service-aws-terraform-registry` ) 

# Regenrate vault-credentials
aws-vault exec sajan.jaiswal --duration=12h


# Terraform workspace

