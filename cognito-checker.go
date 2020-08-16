package gojwtcognito

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
)

// CognitoChecker is the main object of the API of the package.
// Contains information about your AWS Cognito configuration.
type CognitoChecker struct {
	region    string
	userPool  string
	appClient string
	jwks      *jwk.Set
}

// NewCognitoChecker is used for generating a CognitoChecker object and been able to use the library
// Needs the region, user pool id and app client id of your Cognito user pool to work properly
func NewCognitoChecker(region, userPool, appClient string) (*CognitoChecker, error) {
	cognitoChecker := &CognitoChecker{}

	cognitoChecker.region = region
	cognitoChecker.userPool = userPool
	cognitoChecker.appClient = appClient

	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json",
		region,
		userPool)

	jwks, err := jwk.Fetch(jwkURL)
	if err != nil {
		return cognitoChecker, err
	}
	cognitoChecker.jwks = jwks

	return cognitoChecker, nil
}
