package gojwtcognito

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
)

// GetClaims parses a request header and looks for a specific JWT from AWS Cognito.
// Returns a map with all the claims in it or an error if it is an invalid token.
// Use this function when you need the Cognito claims of a token.
func (c CognitoChecker) GetClaims(request *http.Request, tokenType string) (map[string]interface{}, error) {

	encryptedToken, err := c.getJWT(request, tokenType)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(encryptedToken, jwt.WithKeySet(c.jwks))
	if err != nil {
		return nil, fmt.Errorf("error decrypting JWT: %w", err)
	}

	err = c.validateJWT(token.PrivateClaims(), tokenType)
	if err != nil {
		return nil, err
	}

	return token.PrivateClaims(), nil
}
