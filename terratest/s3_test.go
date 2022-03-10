package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestS3Bucket(t *testing.T) {
	t.Parallel()

	tmpTestFolder := test_structure.CopyTerraformFolderToTemp(t, fixtureFolder, "./")

	randomId := strings.ToLower(random.UniqueId())

	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	// kmsKeyArn := CreateKmsKey(t, awsRegion)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tmpTestFolder,
		Vars: map[string]interface{}{
			"bucket_name": fmt.Sprintf("%s-%s", "blah", randomId),
			// "create_bucket": true,
			// "bucket_count":  3,
			// "app":           fmt.Sprintf("%s-%s", app, randomId),
			// "environment":   env,
			// "server_side_encryption_configuration": map[string]interface{}{
			// 	"rule": map[string]interface{}{
			// 		"apply_server_side_encryption_by_default": map[string]interface{}{
			// 			"kms_master_key_id": kmsKeyArn,
			// 			"sse_algorithm":     "aws:kms",
			// 		},
			// 	},
			// },
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
			"AWS_PROFILE":        awsProfile,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	CheckBucketEncryptionEnabled(t, terraformOptions)
}
