package jwtcognito

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
)

// GetJWK is used with in conjunction of the other functions to parse tokens.
// Fetches and parses the JWK in order to build the RSA key needed to decode a JWT
func GetJWK(info CognitoConfig) (*jwk.Set, error) {

	jwkURL := fmt.Sprintf("https://cognito-idp.%v.amazonaws.com/%v/.well-known/jwks.json",
		info.Region,
		info.UserPool)

	jwks, err := jwk.Fetch(jwkURL)
	if err != nil {
		return nil, err
	}

	return jwks, err
}
