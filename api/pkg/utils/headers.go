package utils

import (
	"errors"
	"net/http"
)

var (
	// ErrNoHeader occurs when header does not meet our specification
	ErrNoHeader = errors.New("unknown token")
)

// GetHeaderValue retrieve header value
func GetHeaderValue(header http.Header, name string) (string, error) {
	value := header.Get(name)

	if value == "" || value == "undefined" {
		return "", ErrNoHeader
	}

	return value, nil
}
