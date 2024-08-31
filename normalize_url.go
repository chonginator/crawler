package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error) {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	
	result := parsedURL.Hostname() + parsedURL.Path
	result = strings.ToLower(result)
	result = strings.TrimPrefix(result, "www.")
	result = strings.TrimSuffix(result, "/")

	return result, nil
}