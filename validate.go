package gojwtcognito

import (
	"net/http"
)

// ValidateTokenFromHeader parses a request header and looks for a specific JWT from AWS Cognito.
// Returns an error if it's not valid or nil if it is.
// Use this function when you only need to check if a token is valid or not.
func (c CognitoChecker) ValidateTokenFromHeader(request *http.Request, tokenType string) error {

	_, err := c.GetClaims(request, tokenType)
	if err != nil {
		return err
	}

	return nil
}
