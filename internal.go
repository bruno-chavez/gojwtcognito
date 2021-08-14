// Package gojwtcognito is an easy to use, small package designed to parse request headers and look for
// JWTs provided by AWS Cognito to either check if they are valid or get some data from them.
package gojwtcognito

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"net/http"
	"strings"
)

// fetches from the request header all the necessary info to get the cookie info
func getCookie(request *http.Request, appClient string, tokenType string) (*http.Cookie, error) {

	userIdCookie := fmt.Sprintf("CognitoIdentityServiceProvider.%v.LastAuthUser", appClient)
	userId, err := request.Cookie(userIdCookie)
	if err != nil {
		return nil, fmt.Errorf("could not find cookie with name: %v", userIdCookie)
	}

	tokenName := fmt.Sprintf("CognitoIdentityServiceProvider.%v.%v.%v", appClient, userId.Value, tokenType)
	cookie, err := request.Cookie(tokenName)
	if err != nil {
		return nil, fmt.Errorf("could not find cookie with name: %v", tokenName)
	}

	return cookie, nil
}

// segment is used to unmarshall a part of cookie
type segment struct {
	Kid string `json:"kid"`
}

// getKey builds the key using the jwks and the first part of the specified cookie in order to match key ids
func getKey(cookie *http.Cookie, jwks *jwk.Set) (interface{}, error) {

	// the key id of a JWT is in the first part of the value
	tokenIdSegment := strings.Split(cookie.Value, ".")[0]
	decodedIdSegment, err := jwt.DecodeSegment(tokenIdSegment)
	if err != nil {
		return nil, err
	}

	var seg segment
	err = json.Unmarshal(decodedIdSegment, &seg)
	if err != nil {
		return nil, err
	}

	// matches key id from the cookie and the jwks
	keys := jwks.LookupKeyID(seg.Kid)
	if len(keys) == 0 {
		return nil, fmt.Errorf("failed to look up: %v", seg.Kid)
	}

	key, err := keys[0].Materialize()
	if err != nil {
		return nil, err
	}

	return key, nil
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
