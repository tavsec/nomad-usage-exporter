# Nomad Resources Explorer
This program will fetch Nomad tasks for given time period and export the resource changes to given DynamoDB database.  
The DynamoDB can be created using Terraform script in `./terraform` directory.

## Setup infrastructure
You must have Terraform installed. Then move to `terraform` folder and run:
```bash
terraform apply
echo "DYNAMODB_TABLE_NAME=$(terraform output dynamodb_table_name)" > ../.env
```

It will prefill the .env file with generated DynamoDB name.
