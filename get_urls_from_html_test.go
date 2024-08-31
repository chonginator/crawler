package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name string
		inputBody string
		inputURL string
		expected []string
		errorContains string
	}{
		{
			name: "absolute URL",
			inputBody: `
				<html>
					<body>
						<a href="https://blog.boot.dev">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name: "relative URL",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name: "absolute and relative URLs",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name: "invalid raw base URL",
			inputBody: `
				<html>
					<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
			`,
			inputURL: ":\\invalidURL",
			expected: nil,
			errorContains: "couldn't parse raw base URL",
		},
		{
			name: "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html>
					<body>
						<a>
							<span>Boot.dev></span>
						</a>
					</body>
				</html>
			`,
			expected: nil,
		},
		{
			name: "invalid href URL",
			inputBody: `
				<html>
					<body>
						<a href=":\\invalidURL">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: nil,
			errorContains: "couldn't parse href",
		},
		{
			name: "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
				<html body>
					<a href="path/one">
						<span>Boot.dev></span>
					</a>
				</html body>
			`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name: "nested anchor tags",
			inputBody: `
				<html>
					<body>
						<div>
							<a href="/path/one">
								<span>Boot.dev</span>
							</a>
							<div>
								<a href="/path/two">
									<span>Boot.dev</span>
								</a>
							</div>
						</div>
						<a href="https://other.com/path/one">
							<span>Boot.dev</span>
						</a>
					</body>
				</html>
			`,
			inputURL: "https://blog.boot.dev",
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing: '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}