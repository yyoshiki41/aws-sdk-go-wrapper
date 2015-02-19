// AWS authorization libs

package auth

import (
	AWS "github.com/awslabs/aws-sdk-go/aws"
	"github.com/evalphobia/aws-sdk-go-wrapper/config"
)

const (
	authConfigSectionName = "auth"
	awsAccessConfigKey    = "access_key"
	awsSecretConfigKey    = "secret_key"
)

var (
	auth AWS.CredentialsProvider = nil
)

// return AWS authorization credentials
func Auth() AWS.CredentialsProvider {
	if auth != nil {
		return auth
	}

	// return if environmental params for AWS auth
	env, err := AWS.EnvCreds()
	if err == nil {
		return env
	}

	accessKey := config.GetConfigValue(authConfigSectionName, awsAccessConfigKey, "")
	secretKey := config.GetConfigValue(authConfigSectionName, awsSecretConfigKey, "")
	return AWS.Creds(accessKey, secretKey, "")
}
