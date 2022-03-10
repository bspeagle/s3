<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_s3_bucket.logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket) | resource |
| [aws_s3_bucket.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket) | resource |
| [aws_s3_bucket_policy.redshift](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_policy) | resource |
| [aws_s3_bucket_public_access_block.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_public_access_block) | resource |
| [aws_redshift_service_account.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/redshift_service_account) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_acl"></a> [acl](#input\_acl) | ACL for bucket (public-read, private, etc...) | `string` | `"private"` | no |
| <a name="input_app"></a> [app](#input\_app) | The name of the application | `string` | n/a | yes |
| <a name="input_bucket_count"></a> [bucket\_count](#input\_bucket\_count) | Count of resources to create | `number` | `1` | no |
| <a name="input_bucket_names"></a> [bucket\_names](#input\_bucket\_names) | List o' bucket names | `list(string)` | `[]` | no |
| <a name="input_bucket_type"></a> [bucket\_type](#input\_bucket\_type) | The type of bucket to create | `string` | `"default"` | no |
| <a name="input_create_bucket"></a> [create\_bucket](#input\_create\_bucket) | Should this bucket be created? Should it?!?!?!? | `bool` | `true` | no |
| <a name="input_disable_acls"></a> [disable\_acls](#input\_disable\_acls) | Toggle to disable ACLs for public bucket access | `bool` | `true` | no |
| <a name="input_environment"></a> [environment](#input\_environment) | The application environment | `string` | n/a | yes |
| <a name="input_force_destroy"></a> [force\_destroy](#input\_force\_destroy) | Force destroy all files in bucket when destroying bucket. BE CAREFUL!!! | `bool` | `false` | no |
| <a name="input_kms_key_arn"></a> [kms\_key\_arn](#input\_kms\_key\_arn) | KMS key ARN for server side encryption | `string` | `null` | no |
| <a name="input_logging_enabled"></a> [logging\_enabled](#input\_logging\_enabled) | Toggle bucket logging | `bool` | `false` | no |
| <a name="input_sse_enabled"></a> [sse\_enabled](#input\_sse\_enabled) | Toggle server-side encryption. If true variable kms\_master\_key\_id must be supplied | `bool` | `true` | no |
| <a name="input_versioning_enabled"></a> [versioning\_enabled](#input\_versioning\_enabled) | Toggle bucket versioning | `bool` | `true` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the bucket |
| <a name="output_id"></a> [id](#output\_id) | The Id of the S3 bucket |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
