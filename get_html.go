package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return "", fmt.Errorf("error-level HTTP status: %v", resp.Status)
	}

	mediaType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return "", fmt.Errorf("error parsing media type %s: %v", resp.Header.Get("Content-Type"), err)
	}
	if mediaType != "text/html" {
		return "", fmt.Errorf("content-type header is not text/html: %s", resp.Header.Get("content-type"))
	}

	htmlBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response: %v", err)
	}
	htmlBody := string(htmlBodyBytes)

	return htmlBody, nil
}