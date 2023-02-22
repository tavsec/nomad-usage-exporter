output "dynamodb_table_name" {
    value = aws_dynamodb_table.nomad_resource_table.name
}
