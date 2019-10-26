package gojwtcognito

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"net/http"
)

// GetClaims parses a request header and looks for a specific JWT from AWS Cognito.
// Returns a map with all the claims in it or an error if it is an invalid token.
func GetClaims(request *http.Request, jwks *jwk.Set, info CognitoConfig, tokenType string) (map[string]interface{}, error) {

	cookie, err := getCookie(request, info.AppClient, tokenType)
	if err != nil {
		return nil, err
	}

	key, err := getKey(cookie, jwks)
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
