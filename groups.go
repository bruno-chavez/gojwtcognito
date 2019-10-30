package gojwtcognito

import (
	"github.com/lestrrat-go/jwx/jwk"
	"net/http"
)

// GetGroups parses a request header and looks for a specific JWT from AWS Cognito.
// Returns a slice with all the groups of a user or an error if it is an invalid token.
// Use this function when you need the Cognito groups of a user.
func GetGroups(request *http.Request, jwks *jwk.Set, info CognitoConfig) ([]string, error) {

	claims, err := GetClaims(request, jwks, info, "idToken")
	if err != nil {
		return nil, err
	}

	// Type asserts all the groups to strings since they can't anything else
	var groups []string
	for _, v := range claims["cognito:groups"].([]interface{}) {
		groups = append(groups, v.(string))
	}

	return groups, nil
}
