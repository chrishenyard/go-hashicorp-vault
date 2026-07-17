package gohashicorpvault

import (
	"os"
	"testing"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func TestGetSecrets_ReturnValue(t *testing.T) {
	// Test function to ensure that getSecrets returns a non-nil client and no error.
	options := &Options{
		Address:                       os.Getenv("VAULT_ADDR"),
		AuthMethod:                    "approle",
		KubernetesJwtPath:             os.Getenv("VAULT_KUBERNETES_JWT_PATH"),
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

	var typedSecretList *vault.Response[schema.KvV2ReadResponse] = secretList
	if typedSecretList == nil {
		t.Errorf("Expected return type *vault.Response[schema.KvV2ReadResponse], got %T", secretList)
	}
}
