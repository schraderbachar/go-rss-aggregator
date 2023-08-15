package auth

import (
	"errors"
	"net/http"
	"strings"
)

// * extracts api key from the headers of an http req
// to be used in getUser Handler
// !Example: Authorization: ApiKey {api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no key found")
	}
	//otherwise we have valid value

	vals := strings.Split(val, " ") //split on spaces
	if len(vals) != 2 {
		return "", errors.New("bad auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header (hint: ApiKey)")
	}
	return vals[1], nil
}
