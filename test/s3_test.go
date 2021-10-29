package test

import (
	"fmt"
	// "strings"
	"testing"

	// "github.com/gruntwork-io/terratest/modules/aws"
	// "github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAwsS3Raw(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("test-raw-data-s3")

	// awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",

	})

	defer terraform.Destroy(t, terraformOptions)

    // terraform.InitAndPlan(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)


	bucketID := terraform.Output(t, terraformOptions, "raw_data_s3_id")
	assert.Equal(t, expectedName, bucketID)

	// bucketRegion := terraform.Output(t, terraformOptions, "raw_s3_bucket_region")
	// assert.Equal(t, "us-east-1", bucketRegion)
	
}