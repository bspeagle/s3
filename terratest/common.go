package test

import (
	"fmt"
	"testing"
	"time"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/retry"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var app = "that-darn-app"
var env = "testing"

var fixtureFolder = "../"

var awsProfile = "default-cg"

func CreateKmsKey(t *testing.T, awsRegion string) string {
	kmsClient := aws.NewKmsClient(t, awsRegion)

	kmsInput := &kms.CreateKeyInput{}
	kmsKey, err := kmsClient.CreateKey(kmsInput)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	kmsKeyArn := aws.GetCmkArn(t, awsRegion, *kmsKey.KeyMetadata.KeyId)

	kmsDeleteSchedule := &kms.ScheduleKeyDeletionInput{}
	kmsDeleteSchedule.KeyId = kmsKey.KeyMetadata.KeyId
	kmsDeleteSchedule.PendingWindowInDays = awssdk.Int64(7)
	defer kmsClient.ScheduleKeyDeletion(kmsDeleteSchedule)

	return kmsKeyArn
}

func CheckBucketPolicy(t *testing.T, terraformOptions *terraform.Options) {
	aws.AssertS3BucketPolicyExists(t, terraformOptions.EnvVars["AWS_DEFAULT_REGION"], terraformOptions.Vars["bucket_name"].(string))
}

func CheckBucketEncryptionEnabled(t *testing.T, terraformOptions *terraform.Options) {
	awsRegion := terraformOptions.EnvVars["AWS_DEFAULT_REGION"]

	s3Client, err := aws.NewS3ClientE(t, awsRegion)
	if err != nil {
		assert.FailNow(t, "Error creating s3client")
		return
	}

	bucketName := fmt.Sprintf("%s-%s-1", terraformOptions.Vars["app"].(string), terraformOptions.Vars["environment"].(string))

	params := &s3.GetBucketEncryptionInput{
		Bucket: awssdk.String(bucketName),
	}

	maxRetries := 3
	retryDuration := time.Duration(30)
	_, retryErr := retry.DoWithRetryE(t, "Get bucket encryption", maxRetries, retryDuration,
		func() (string, error) {
			encryption, err := s3Client.GetBucketEncryption(params)

			if err != nil {
				assert.FailNow(t, "Error getting bucket encryption")
				return "", nil
			}

			encryption.ServerSideEncryptionConfiguration.Validate()

			return "Retrieved bucket encryption", nil
		},
	)
	if retryErr != nil {
		assert.FailNow(t, "Error on retry")
		return
	}
}
