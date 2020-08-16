package gojwtcognito

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
)

// Cognito is used for passing necessary information to the API of the package.
// Contains information about your AWS Cognito configuration.
type CognitoChecker struct {
	region    string
	userPool  string
	appClient string
	jwks      *jwk.Set
}

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
