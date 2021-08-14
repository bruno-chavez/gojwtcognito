// Package gojwtcognito is an easy-to-use, small wrapper over 'github.com/lestrrat-go/jwx/jwk'
// to provide specific functionality to work with AWS Cognito JWTs
package gojwtcognito

import (
	"fmt"
	"net/http"
)

// Parses a http request to get the JWT
func (c CognitoChecker) getJWT(request *http.Request, tokenType string) ([]byte, error) {

	lastAuthUserKey := fmt.Sprintf("CognitoIdentityServiceProvider.%v.LastAuthUser", c.appClient)
	userId, err := request.Cookie(lastAuthUserKey)
	if err != nil {
		return nil, fmt.Errorf("could not find 'LastAuthUser' cookie with name: %v", lastAuthUserKey)
	}

	tokenName := fmt.Sprintf("CognitoIdentityServiceProvider.%v.%v.%v", c.appClient, userId.Value, tokenType)
	cookie, err := request.Cookie(tokenName)
	if err != nil {
		return nil, fmt.Errorf("could not find '%v' cookie with name: %v", tokenType, tokenName)
	}

	return []byte(cookie.Value), nil
}

func (c CognitoChecker) validateJWT(claims map[string]interface{}, tokenType string) error {

	// verifies token_use, issuer/client_id and target audience inside a JWT
	switch {
	case tokenType == "idToken":
		if claims["token_use"] != "id" {
			return fmt.Errorf("invalid token use: %v", claims["token_use"])
		}
		if claims["aud"] != c.appClient {
			return fmt.Errorf("invalid target audience: %v", claims["aud"])
		}
	case tokenType == "accessToken":
		if claims["token_use"] != "access" {
			return fmt.Errorf("invalid token use: %v", claims["token_use"])
		}
		if claims["client_id"] != c.appClient {
			return fmt.Errorf("invalid target audience: %v", claims["client_id"])
		}
	}

	iss := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v", c.region, c.userPool)
	if claims["iss"] != iss {
		return fmt.Errorf("invalid issuer: %v", claims["iss"])
	}

	return nil
}
