[![GoDoc](https://godoc.org/github.com/bruno-chavez/gojwtcognito?status.svg)](https://godoc.org/github.com/bruno-chavez/gojwtcognito)
[![Build Status](https://travis-ci.org/bruno-chavez/gojwtcognito.svg?branch=master)](https://travis-ci.org/bruno-chavez/gojwtcognito)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruno-chavez/gojwtcognito)](https://goreportcard.com/report/github.com/bruno-chavez/gojwtcognito)

`gojwtcognito` is an easy to use, small package 
designed to parse request headers
and look for JWTs provided by AWS Cognito 
to either check if they are valid or get info from them.

##  Install

```
$ go get github.com/bruno-chavez/gojwtcognito
```

## Usage

Import the package, call a `NewCognitoChecker`.
From here you pass the object pointer to where you need to
validate tokens, get claims or groups.

### Usage Tips 

+ The only two tokens that the library works with are `idToken` and `accessToken`.

+ Note that the username password (`ALLOW_USER_PASSWORD_AUTH`)
 based authentication flow is not supported.

+ The region, User Pool ID and App Client 
can all be found inside AWS Cognito.

+ The claims inside each JWT varies depends on the token type
you pass to `GetClaims`. Please check this [link](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-with-identity-providers.html) 
for the official specification and usage of each token type.

## Documentation

Check the [GoDoc](https://godoc.org/github.com/bruno-chavez/gojwtcognito)
page for more info on what is available inside the package.

## Examples

### Generating the checker object
```go

checker := gojwtcognito.NewCognitoChecker(
                       "us-east-1", // region
                       "us-east-1_apwePSzx", // user pool id
                       "3b1fh12qzvmgjuio563qtm678u", // client app id
                  )
```

### Validating an accessToken
```go
func ExampleHandler(w http.ResponseWriter, r *http.Request) {

    err := checker.ValidateTokenFromHeader(r, "accessToken")
    if err != nil {
        log.Println(err)
        return
    }
    
    err = c.ValidateTokenFromHeader(r, "idToken")
    if err != nil {
        log.Println(err)
        return
    }
}
```

#### Looking up a specific claim
`claims` is a map of type `map[string]interface{}`
```go
func ExampleHandler(w http.ResponseWriter, r *http.Request) {

    claims, err := c.GetClaims(r, "idToken")
    if err != nil {
        log.Println(err)
    }
    
    log.Println(claims["cognito:username"])
}
```

#### Looking up all the groups of a user
`groups` is a slice of type `[]string`
```go
func ExampleHandler(w http.ResponseWriter, r *http.Request) {

    groups, err := checker.GetGroups(r)
    if err != nil {
        log.Println(err)
    }
    
    for _, v := range groups {
        fmt.Println(v)
    }
}
```
 

## Contribute

Found a bug or an error? Post it in the 
[issue tracker](https://github.com/bruno-chavez/gojwtcognito/issues).

Want to add an awesome new feature? 
[Fork](https://github.com/bruno-chavez/gojwtcognito/fork) 
this repository and add your feature, then send a pull request.

## License
The MIT License (MIT)
Copyright (c) 2020 Bruno Chavez
