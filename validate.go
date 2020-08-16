package gojwtcognito

import (
	"fmt"
	"net/http"
)

// ValidateTokenFromHeader parses a request header and looks for a specific JWT from AWS Cognito.
// Returns an error if its not valid or nil if it is.
// Use this function when you only need to check if a token is valid or not.
func (c CognitoChecker) ValidateTokenFromHeader(request *http.Request, tokenType string) error {

	claims, err := c.GetClaims(request, tokenType)
	if err != nil {
		return err
	}

	// verifies token_use, issuer and target audience inside a JWT
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
	}

	iss := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v", c.region, c.userPool)
	if claims["iss"] != iss {
		return fmt.Errorf("invalid issuer: %v", claims["iss"])
	}

	return nil
}
