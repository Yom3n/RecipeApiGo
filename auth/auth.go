package auth

import (
	"errors"
	"net/http"
	"strings"
)

type ApiKey string

// Gets api key headers. Correct header format is "Authorization: ApiKey Value"
func GetApiKey(header http.Header) (ApiKey, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not provided")
	}

	splitHeader := strings.Split(authHeader, " ")
	if len(splitHeader) != 2 {
		return "", errors.New("wrong format of authorization header")
	}
	if splitHeader[0] != "ApiKey" {
		return "", errors.New("malformed first part of authorization header")
	}

	apiKey := splitHeader[1]
	return ApiKey(apiKey), nil

}
