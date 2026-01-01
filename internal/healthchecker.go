package internal

import (
	"errors"
	"fmt"
	"net/url"
)

// HealthChecker is the interface for structs that implement health checks
type HealthChecker interface {

	// Name should return a constant, human-friendly name common to all instances of an implementation
	// (e.g. "http" for the HTTP and HTTPs implementation).
	Name() string

	// Check performs the actual health check and returns an error in case of failure
	Check() error
}

// GetHealthCheckerForUrl takes url as a string and returns the appropriate HealthChecker for that URL.
// This function fails early: only a valid URL shall return a HealthChecker and with no error.
func GetHealthCheckerForUrl(urlString string) (HealthChecker, error) {
	if urlString == "" {
		return nil, errors.New("no URL to check was supplied, provide a valid one at build-time")
	}
	urlParsed, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	if urlParsed.Scheme == "" {
		return nil, fmt.Errorf("missing scheme in url: '%s'", urlString)
	}
	switch urlParsed.Scheme {
	case "http", "https":
		return &HttpHealthChecker{url: *urlParsed}, nil
	default:
		return nil, fmt.Errorf("unsupported scheme in url: '%s'", urlParsed.Scheme)
	}
}
