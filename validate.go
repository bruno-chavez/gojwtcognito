package jwtcognito

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"net/http"
)

// ValidateTokenFromHeader parses a request header and looks for a specific JWT from AWS Cognito.
// Returns an error if its not valid or nil if it is.
func ValidateTokenFromHeader(request *http.Request, jwks *jwk.Set, info CognitoConfig, tokenType string) error {

	claims, err := GetClaims(request, jwks, info, tokenType)
	if err != nil {
		return err
	}

	// verifies token_use, issuer and target audience inside a JWT
	switch {
	case tokenType == "idToken":
		if claims["token_use"] != "id" {
			return fmt.Errorf("invalid token use: %v", claims["token_use"])
		}

		if claims["aud"] != info.AppClient {
			return fmt.Errorf("invalid target audience: %v", claims["aud"])
		}
	case tokenType == "accessToken":
		if claims["token_use"] != "access" {
			return fmt.Errorf("invalid token use: %v", claims["token_use"])
		}
	}

	iss := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v", info.Region, info.UserPool)
	if claims["iss"] != iss {
		return fmt.Errorf("invalid issuer: %v", claims["iss"])
	}

	return nil
}
