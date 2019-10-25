package jwtcognito

// CognitoConfig is used for passing necessary information to the API of the package.
// Contains information about your AWS Cognito configuration.
type CognitoConfig struct {
	Region    string
	UserPool  string
	AppClient string
}
