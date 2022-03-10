resource "aws_s3_bucket_policy" "redshift" {
  count  = var.bucket_type == "redshift" ? var.create_bucket == true ? 1 : 0 : 0
  bucket = aws_s3_bucket.this[count.index].id

  policy = jsonencode({
    "Version" : "2008-10-17",
    "Id" : var.bucket_type == "iics" ? var.bucket_names[count.index] : "${local.bucket_name_prefix}-${var.bucket_names[count.index]}"
    "Statement" : [
      {
        "Sid" : "Put bucket policy needed for audit logging",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "${data.aws_redshift_service_account.this[count.index].arn}"
        },
        "Action" : "s3:PutObject",
        "Resource" : "arn:aws:s3:::${aws_s3_bucket.this[count.index].id}/*"
      },
      {
        "Sid" : "Get bucket policy needed for audit logging ",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "${data.aws_redshift_service_account.this[count.index].arn}"
        },
        "Action" : "s3:GetBucketAcl",
        "Resource" : "arn:aws:s3:::${aws_s3_bucket.this[count.index].id}"
      }
    ]
  })
}
