package internal

import (
	"errors"
	"fmt"
	"net/url"
)

type HealthChecker interface {
	Name() string
	Check() error
}

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
