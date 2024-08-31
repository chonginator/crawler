package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name string
		inputURL string
		expected string
		errorContains string
	}{
		{
			name: "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name: "remove www",
			inputURL: "https://www.boot.dev/path",
			expected: "boot.dev/path",
		},
		{
			name: "remove user info",
			inputURL: "https://user:password@boot.dev/path",
			expected: "boot.dev/path",
		},
		{
			name: "remove port",
			inputURL: "https://boot.dev:443/path",
			expected: "boot.dev/path",
		},
		{
			name: "remove query parameters",
			inputURL: "https://google.com/search?q=boot.dev",
			expected: "google.com/search",
		},
		{
			name: "remove hash fragment",
			inputURL: "https://blog.boot.dev/path#fragment",
			expected: "blog.boot.dev/path",
		},
		{
			name: "lowercase capital letters",
			inputURL: "https://blog.BOOT.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name: "remove scheme, trailing slash, www, user info, port, query parameters, hash fragment, and capital letters",
			inputURL: "https://user:password@www.BOOT.dev:443/path?query=value#fragment",
			expected: "boot.dev/path",
		},
		{
			name: "handle invalid URL",
			inputURL: ":\\invalidURL",
			expected: "",
			errorContains: "couldn't parse URL",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func (t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)

			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing: '%v', got none.", i, tc.name, tc.errorContains)
			}

			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}