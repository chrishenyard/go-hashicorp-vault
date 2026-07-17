package gohashicorpvault

import (
	"context"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func getSecrets(options *Options) (*vault.Response[schema.KvV2ReadResponse], error) {
	ctx := context.Background()

	client, err := vault.New(
		vault.WithAddress(options.Address),
		vault.WithRequestTimeout(30*time.Second),
		vault.WithTLS(vault.TLSConfiguration{
			InsecureSkipVerify: options.AllowInvalidServerCertificate,
		}),
	)
	if err != nil {
		return nil, err
	}

	resp, err := authenticateWithVault(ctx, client, options)
	if err != nil {
		return nil, err
	}

	if err := client.SetToken(resp.Auth.ClientToken); err != nil {
		return nil, err
	}

	secretList, err := client.Secrets.KvV2Read(
		ctx,
		options.SecretPath,
	)
	if err != nil {
		return nil, err
	}

	return secretList, nil
}

func authenticateWithVault(ctx context.Context, client *vault.Client, options *Options) (*vault.Response[map[string]interface{}], error) {
	switch options.AuthMethod {
	case "approle":
		resp, err := client.Auth.AppRoleLogin(
			ctx,
			schema.AppRoleLoginRequest{
				RoleId:   options.RoleId,
				SecretId: options.SecretId,
			},
		)
		if err != nil {
			return nil, err
		}

		if err := client.SetToken(resp.Auth.ClientToken); err != nil {
			return nil, err
		}

		return resp, nil
	case "kubernetes":
		jwt, err := os.ReadFile(options.KubernetesJwtPath)
		if err != nil {
			return nil, err
		}

		resp, err := client.Auth.KubernetesLogin(
			ctx,
			schema.KubernetesLoginRequest{
				Role: options.RoleName,
				Jwt:  string(jwt),
			},
		)
		if err != nil {
			return nil, err
		}

		if err := client.SetToken(resp.Auth.ClientToken); err != nil {
			return nil, err
		}

		return resp, nil
	default:
		return nil, nil
	}
}
