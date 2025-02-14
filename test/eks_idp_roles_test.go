package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestEksIdpRoles(t *testing.T) {
	t.Parallel()

	// Root folder where terraform files should be (up one level from /test)
	terraformDir := "../src"

	// Terraform options for testing
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: terraformDir,
		// Set any variables needed for testing
		Vars: map[string]interface{}{
			// Add any required variables here
		},
	})

	// Init terraform
	terraform.Init(t, terraformOptions)

	// The rest of this test is intentionally left blank
	// We are still migrating tests to components and this is a placeholder
}