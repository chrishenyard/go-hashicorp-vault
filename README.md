# go-hashicorp-vault

A lightweight Go package for reading secrets from HashiCorp Vault KV v2.

The package currently supports two Vault authentication flows:

- `approle`
- `kubernetes`

## Status

This repository is in an early stage. The core secret retrieval logic exists and is covered by a basic test, but the primary retrieval function is not currently exported.

## Requirements

- Go `1.26.5` (as defined in `go.mod`)
- A reachable HashiCorp Vault instance
- A KV v2 secret path in Vault

## Configuration

The package uses an `Options` struct:

```go
type Options struct {
	Provider                      string
	Address                       string
	HostHeader                    string
	AuthMethod                    string
	Token                         string
	RoleName                      string
	KubernetesJwtPath             string
	RoleId                        string
	SecretId                      string
	MountPoint                    string
	SecretPath                    string
	AllowInvalidServerCertificate bool
}
```

### Fields used by current implementation

- `Address`
- `AuthMethod` (`approle` or `kubernetes`)
- `RoleId` and `SecretId` (for `approle`)
- `RoleName` and `KubernetesJwtPath` (for `kubernetes`)
- `SecretPath`
- `AllowInvalidServerCertificate`

Note: `MountPoint`, `Provider`, `HostHeader`, and `Token` are present in the struct but are not currently used by the active retrieval path.

## Authentication

### AppRole

Uses Vault AppRole login with:

- `RoleId`
- `SecretId`

### Kubernetes

Reads a JWT from `KubernetesJwtPath` and performs Vault Kubernetes auth with:

- `RoleName`
- JWT file contents

## Testing

The included test (`gohashicorpvault_test.go`) expects Vault-related environment variables to be set:

- `VAULT_ADDR`
- `VAULT_KUBERNETES_JWT_PATH`
- `VAULT_ROLE_ID`
- `VAULT_SECRET_ID`
- `VAULT_MOUNT_POINT`
- `VAULT_SECRET_PATH`

Run tests with:

```bash
go test ./...
```

## Notes

- The client is configured with a 30-second request timeout.
- TLS verification can be disabled via `AllowInvalidServerCertificate` (use only in trusted/non-production environments).
