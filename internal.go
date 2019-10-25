// Package jwtcognito is an easy to use, small package designed to parse request headers and look for
// JWTs provided by AWS Cognito to either check if they are valid or get some data from them.
package jwtcognito

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"net/http"
	"strings"
)

// fetches from the request header all the necessary info to get the cookie where the JWT lays
func getCookie(request *http.Request, appClient string, tokenType string) (*http.Cookie, error) {

	appClientCookieName := fmt.Sprintf("CognitoIdentityServiceProvider.%v.LastAuthUser", appClient)
	userId, err := request.Cookie(appClientCookieName)
	if err != nil {
		return nil, err
	}

	tokenName := fmt.Sprintf("CognitoIdentityServiceProvider.%v.%v.%v", appClient, userId.Value, tokenType)
	cookie, err := request.Cookie(tokenName)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}

// segment is used to unmarshall a part of cookie
type segment struct {
	Kid string `json:"kid"`
}

// getkey builds the key using the JWKs and the first part of the specified cookie in order to match key ids
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

	// matches key id from the cookie and the JWKs
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

// parses a JWT using the cookie where it is contained and the key obtained from getKey() and returns it.
func getToken(cookie *http.Cookie, key interface{}) (*jwt.Token, error) {

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
