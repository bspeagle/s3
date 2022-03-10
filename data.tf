data "aws_redshift_service_account" "this" {
  count = var.create_bucket == true ? 1 : 0
}