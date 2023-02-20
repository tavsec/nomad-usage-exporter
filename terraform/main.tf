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
    billing_mode     = "PAY_PER_REQUEST"
    hash_key         = "ID"

    attribute {
        name = "ID"
        type = "S"
    }

    attribute {
        name = "TaskName"
        type = "S"
    }

    attribute {
        name = "VersionId"
        type = "N"
    }

    global_secondary_index {
        name               = "VersionId-Index"
        hash_key           = "VersionId"
        write_capacity     = 1
        read_capacity      = 1
        projection_type    = "ALL"
    }
    global_secondary_index {
        name               = "TaskName-Index"
        hash_key           = "TaskName"
        write_capacity     = 1
        read_capacity      = 1
        projection_type    = "ALL"
    }

    tags = {
        Name        = "nomad-resources-table"
    }
}
