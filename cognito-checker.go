package gojwtcognito

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
)

// CognitoChecker is the main object of the API of the package.
// Contains information about your AWS Cognito configuration.
type CognitoChecker struct {
	region    string
	userPool  string
	appClient string
	jwks      jwk.Set
}

// NewCognitoChecker is used for generating a CognitoChecker object and been able to use the library.
// Needs the region, user pool id and app client id of your Cognito user pool to work.
func NewCognitoChecker(region, userPool, appClient string) (*CognitoChecker, error) {
	cognitoChecker := &CognitoChecker{
		region:    region,
		userPool:  userPool,
		appClient: appClient,
	}

	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json",
		region,
		userPool)

	jwks, err := jwk.Fetch(context.TODO(), jwkURL)
	if err != nil {
		return cognitoChecker, fmt.Errorf("error fetching jwks: %w", err)
	}
	cognitoChecker.jwks = jwks

	return cognitoChecker, nil
}
