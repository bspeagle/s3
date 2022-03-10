locals {
  bucket_name_prefix = "${var.tags.App}-${var.tags.Env}"

  logging_configuration = {
    target_bucket = var.logging_enabled == true ? aws_s3_bucket.logging[0].bucket : ""
    target_prefix = "log/"
  }

  server_side_encryption_configuration = {
    rule = {
      apply_server_side_encryption_by_default = {
        kms_master_key_id = var.sse_enabled == true ? var.kms_key_arn : ""
        sse_algorithm     = "aws:kms"
      }
    }
  }
}
