package gojwtcognito

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

// GetClaims parses a request header and looks for a specific JWT from AWS Cognito.
// Returns a map with all the claims in it or an error if it is an invalid token.
// Use this function when you need the Cognito claims of a token.
func (c CognitoChecker) GetClaims(request *http.Request, tokenType string) (map[string]interface{}, error) {

	cookie, err := getCookie(request, c.appClient, tokenType)
	if err != nil {
		return nil, err
	}

	key, err := getKey(cookie, c.jwks)
	if err != nil {
		return nil, err
	}

	token, err := getToken(cookie, key)
	if err != nil {
		return nil, err
	}

	// goes over all the claims and adds them to a map
	claims := make(map[string]interface{})
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	for i, v := range mapClaims {
		claims[i] = v
	}

	return claims, nil
}
