package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAwsS3Raw(t *testing.T) {
	t.Parallel()

	expectedName := fmt.Sprintf("test-raw-data-s3")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",

	})

	defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)


	bucketID := terraform.Output(t, terraformOptions, "raw_data_s3_id")
	assert.Equal(t, expectedName, bucketID)
	
}
