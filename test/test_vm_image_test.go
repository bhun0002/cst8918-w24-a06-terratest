package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/ssh"
	"github.com/stretchr/testify/assert"
)

func TestAzureVMUbuntuVersion(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Get the VM Public IP from Terraform output
	publicIP := terraform.Output(t, terraformOptions, "public_ip")

	// Define the SSH key and user
	sshUser := "azureadmin"
	sshKey := "~/.ssh/id_rsa"

	sshHost := ssh.Host{
		Hostname:    publicIP,
		SshKeyPair:  &ssh.KeyPair{PrivateKeyPath: sshKey},
		SshUserName: sshUser,
	}

	// Run the command to check Ubuntu version
	ubuntuVersion, err := ssh.CheckSshCommandE(t, sshHost, "lsb_release -d | awk -F':' '{print $2}'")
	assert.NoError(t, err)
	assert.Contains(t, ubuntuVersion, "Ubuntu 22.04", "Expected Ubuntu 22.04 LTS")
}
