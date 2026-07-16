package gohashicorpvault

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
