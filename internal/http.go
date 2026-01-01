package internal

import (
	"fmt"
	"net/http"
	"net/url"
)

// HttpHealthChecker is the implementation of HealthChecker for HTTP and HTTPS URLs.
type HttpHealthChecker struct {
	url url.URL
}

func (h HttpHealthChecker) Name() string {
	return "http"
}

func (h HttpHealthChecker) Check() error {
	client := http.Client{}
	response, err := client.Head(h.url.String())
	if err != nil {
		return err
	}
	_ = response.Body.Close() // We don't actually need the body
	switch response.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("server did not return a success: %s (%d)", response.Status, response.StatusCode)
	}
}
