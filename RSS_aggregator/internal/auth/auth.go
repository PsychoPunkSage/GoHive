package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")
var ErrMalformedHeader = errors.New("malformed authorization header")

// GetAPIKey -
// Extract API key from the `header` of an HTTP request.
func GetAPIKey(headers http.Header) (string, error) {
	/*
		Expected Header:
			- Authorization: ApiKey {api_key}
	*/
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", ErrMalformedHeader
	}

	return splitAuth[1], nil
}
