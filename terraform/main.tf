terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

resource "aws_dynamodb_table" "nomad_resource_table" {
    name             = "NomadResources"
    billing_mode     = "PROVISIONED"
    read_capacity    = 1
    write_capacity   = 1
    hash_key         = "TaskName"

    attribute {
        name = "TaskName"
        type = "S"
    }

    attribute {
        name = "Timestamp"
        type = "N"
    }

    global_secondary_index {
        name               = "Timestamp-Index"
        hash_key           = "Timestamp"
        write_capacity     = 1
        read_capacity      = 1
        projection_type    = "INCLUDE"
        non_key_attributes = ["Genre"]
    }
    tags = {
        Name        = "nomad-resources-table"
    }
}
