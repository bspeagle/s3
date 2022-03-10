resource "aws_s3_bucket" "logging" {
  count = var.logging_enabled == true ? 1 : 0

  bucket = "${local.bucket_name_prefix}-logging"
  acl    = "log-delivery-write"
}

resource "aws_s3_bucket" "this" {
  count = var.create_bucket == true ? length(var.bucket_names) : 0

  bucket        = var.bucket_type == "iics" ? var.bucket_names[count.index] : "${local.bucket_name_prefix}-${var.bucket_names[count.index]}"
  force_destroy = var.force_destroy
  acl           = var.acl

  versioning {
    enabled = var.versioning_enabled
  }

  dynamic "logging" {
    for_each = var.logging_enabled == true ? [local.logging_configuration] : []

    content {
      target_bucket = local.logging_configuration.target_bucket
      target_prefix = "log/"
    }
  }

  dynamic "server_side_encryption_configuration" {
    for_each = var.sse_enabled == true ? [local.server_side_encryption_configuration] : []

    content {
      dynamic "rule" {
        for_each = [lookup(server_side_encryption_configuration.value, "rule", {})]

        content {
          dynamic "apply_server_side_encryption_by_default" {
            for_each = [lookup(rule.value, "apply_server_side_encryption_by_default", {})]

            content {
              sse_algorithm     = apply_server_side_encryption_by_default.value.sse_algorithm
              kms_master_key_id = apply_server_side_encryption_by_default.value.kms_master_key_id
            }
          }
        }
      }
    }
  }

  tags = merge(var.tags, {
    Name = var.bucket_type == "iics" ? var.bucket_names[count.index] : "${local.bucket_name_prefix}-${var.bucket_names[count.index]}"
  })
}

resource "aws_s3_bucket_public_access_block" "this" {
  count = var.disable_acls == true ? length(var.bucket_names) : 0

  bucket = aws_s3_bucket.this[count.index].id

  block_public_acls   = true
  block_public_policy = true
}
