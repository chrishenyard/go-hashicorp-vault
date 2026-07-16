package gohashicorpvault

import (
	"context"
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

	secretList, err := client.Secrets.KvV2Read(
		ctx,
		options.SecretPath,
	)
	if err != nil {
		return nil, err
	}

	return secretList, nil
}
