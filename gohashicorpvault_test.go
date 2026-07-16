package gohashicorpvault

import (
	"os"
	"testing"
)

func TestGetSecrets_ReturnValue(t *testing.T) {
	// Test function to ensure that getSecrets returns a non-nil client and no error.
	options := &Options{
		Address:                       os.Getenv("VAULT_ADDR"),
		RoleId:                        os.Getenv("VAULT_ROLE_ID"),
		SecretId:                      os.Getenv("VAULT_SECRET_ID"),
		MountPoint:                    os.Getenv("VAULT_MOUNT_POINT"),
		SecretPath:                    os.Getenv("VAULT_SECRET_PATH"),
		AllowInvalidServerCertificate: true,
	}

	secretList, err := getSecrets(options)
	if secretList == nil || err != nil {
		t.Errorf("Expected non-nil secretList and no error, got secretList: %v, err: %v", secretList, err)
	}
}
