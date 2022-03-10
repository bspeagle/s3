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

func TestS3RedShiftBucket(t *testing.T) {
	t.Parallel()

	tmpTestFolder := test_structure.CopyTerraformFolderToTemp(t, fixtureFolder, "./")

	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	bucketName := fmt.Sprintf("test-bucket-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: tmpTestFolder,
		Vars: map[string]interface{}{
			"create_bucket": true,
			"bucket_name":   bucketName,
			"bucket_type":   "redshift",
			"force_destroy": true,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
			"AWS_PROFILE":        awsProfile,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	CheckBucketPolicy(t, terraformOptions)
}
