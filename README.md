[![GoDoc](https://godoc.org/github.com/bruno-chavez/gojwtcognito?status.svg)](https://godoc.org/github.com/bruno-chavez/gojwtcognito)
[![Build Status](https://travis-ci.org/bruno-chavez/gojwtcognito.svg?branch=master)](https://travis-ci.org/bruno-chavez/gojwtcognito)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruno-chavez/gojwtcognito)](https://goreportcard.com/report/github.com/bruno-chavez/gojwtcognito)

`gojwtcognito` is an easy to use, small package 
designed to parse request headers
and look for JWTs provided by AWS Cognito 
to either check if they are valid or get the info in them.

##  Install

```
$ go get github.com/bruno-chavez/gojwtcognito
```

## Usage

Import the package, build a `CognitoConfig` struct and 
get the set of JWKs with `GetJWK`. From here you can 
validate tokens, get claims or groups, depending on your needs.

### Usage Tips 

+ You only need to build the `CognitoConfig` struct and 
call the `GetJWK` function ONCE, 
you can then pass them to other parts of your code 
that get called multiple times, 
like server handlers for example.

+ The region, User Pool ID and App Client 
can all be found inside AWS Cognito.

+ The supported tokens that Cognito uses for 
authentication are either ```accessToken``` or  ```idToken```.

+ If you need to check for user data like emails, 
names or associated groups
```idToken``` is what you are looking for.

+ The claims inside each JWT varies 
depends of the token type.


## Documentation

Check the [GoDoc](https://godoc.org/github.com/bruno-chavez/gojwtcognito)
page for more info on what is available inside the package.

## Examples

#### Using CognitoConfig and GetJWK
You can for example set the CognitoConfig struct 
and call GetJWK globally inside a package 
and then use them inside server handlers.
```

var cognitoConfig = gojwtcognito.CognitoConfig{
                       Region: "us-east-1",
                       UserPool: "us-east-1_apwePSzx",
                       Appclient: "3b1fh12qzvmgjuio563qtm678u",
                     }

var jwks, _ = gojwtcognito.GetJWK(cognitoConfig)

func verifyUser(w http.ResponseWriter, r *http.Request) {

    err := gojwtcognito.ValidateTokenFromHeader(r, jwks, cognitoConfig, "accessToken")
    if err != nil {
        log.Println(err)
        return
    }
    
    err = gojwtcognito.ValidateTokenFromHeader(r, jwks, cognitoConfig, "idToken")
    if err != nil {
        log.Println(err)
        return
    }

}
```

#### Validating an accessToken

```

cognitoConfig := gojwtcognito.CognitoConfig{
    Region: "us-east-1",
    UserPool: "us-east-1_apwePSzx",
    Appclient: "3b1fh12qzvmgjuio563qtm678u",
}

jwks, err := gojwtcognito.GetJWK(cognitoConfig)
if err != nil {
    log.Println(err)
}

err := gojwtcognito.ValidateTokenFromHeader(r, jwks, cognitoConfig, "accessToken")
if err != nil {
    log.Println(err)
}

```

#### Looking up a specific claim
`claims` is a map of type `map[string]interface{}`
```
cognitoConfig := gojwtcognito.CognitoConfig{
    Region: "us-east-1",
    UserPool: "us-east-1_apwePSzx",
    Appclient: "3b1fh12qzvmgjuio563qtm678u",
}

jwks, err := gojwtcognito.GetJWK(cognitoConfig)
if err != nil {
    log.Println(err)
}

claims, err := gojwtcognito.GetClaims(r, jwks, cognitoConfig, "idToken")
if err != nil {
    log.Println(err)
}

log.Println(claims["cognito:username"])
```

#### Looking up all the groups of a user
`groups` is a slice of type `[]string`
```
cognitoConfig := gojwtcognito.CognitoConfig{
    Region: "us-east-1",
    UserPool: "us-east-1_apwePSzx",
    Appclient: "3b1fh12qzvmgjuio563qtm678u",
}

jwks, err := gojwtcognito.GetJWK(cognitoConfig)
if err != nil {
    log.Println(err)
}

groups, err := gojwtcognito.GetGroups(r, jwks, cognitoConfig)
if err != nil {
    log.Println(err)
}

for _, v := range groups {
    fmt.Println(v)
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
Copyright (c) 2019 Bruno Chavez
