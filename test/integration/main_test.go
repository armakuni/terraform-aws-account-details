package integration_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewS3Client() (*s3.Client, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(awsConfig), nil
}

func GetCanonicalUserID(s3Client *s3.Client) (*string, error) {
	response, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return response.Owner.ID, nil
}

func TestAccountDetails(t *testing.T) {
	s3Client, err := NewS3Client()
	require.Nil(t, err, "failed create S3 client")

	options := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/complete",
	})

	defer terraform.Destroy(t, options)
	terraform.InitAndApply(t, options)

	canonicalUserIdOutput := terraform.OutputMapOfObjects(t, options, "canonical_user_id")

	expected, err := GetCanonicalUserID(s3Client)
	require.Nil(t, err, "failed to get the canonical user ID")

	assert.Equal(t, *expected, canonicalUserIdOutput["id"])
}
