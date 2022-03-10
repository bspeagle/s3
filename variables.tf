#--------------------------------------------------------------
# General
#--------------------------------------------------------------

variable "tags" {
  description = "A map of tags to add to all resources"
  type        = map(string)
  default     = {}
}

#--------------------------------------------------------------
# S3
#--------------------------------------------------------------

variable "bucket_names" {
  description = "List o' bucket names"
  type        = list(string)
  default     = []
}

variable "create_bucket" {
  description = "Should this bucket be created? Should it?!?!?!?"
  type        = bool
  default     = true
}

variable "bucket_count" {
  description = "Count of resources to create"
  type        = number
  default     = 1
}

variable "bucket_type" {
  description = "The type of bucket to create"
  type        = string
  default     = "default"
}

variable "force_destroy" {
  description = "Force destroy all files in bucket when destroying bucket. BE CAREFUL!!!"
  type        = bool
  default     = false
}

variable "disable_acls" {
  description = "Toggle to disable ACLs for public bucket access"
  type        = bool
  default     = true
}

variable "acl" {
  description = "ACL for bucket (public-read, private, etc...)"
  type        = string
  default     = "private"
}

variable "versioning_enabled" {
  description = "Toggle bucket versioning"
  type        = bool
  default     = true
}

variable "logging_enabled" {
  description = "Toggle bucket logging"
  type        = bool
  default     = false
}

variable "sse_enabled" {
  description = "Toggle server-side encryption. If true variable kms_master_key_id must be supplied"
  type        = bool
  default     = true
}

variable "kms_key_arn" {
  description = "KMS key ARN for server side encryption"
  type        = string
  default     = null
}
