package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct{
		name string
		inputPages map[string]int
		expected []pageCount
		errorContains string
	}{
		{
			name: "order count descending",
			inputPages: map[string]int{
				"https://example.com/page2": 2,
				"https://example.com/page3": 1,
				"https://example.com/page1": 3,
			},
			expected: []pageCount{
				{
					normalizedURL: "https://example.com/page1",
					count: 3,
				},
				{
					normalizedURL: "https://example.com/page2",
					count: 2,
				},
				{
					normalizedURL: "https://example.com/page3",
					count: 1,
				},
			},
		},
		{
			name: "order count descending then alphabetical",
			inputPages: map[string]int{
				"https://example.com/page4": 2,
				"https://example.com/page2": 2,
				"https://example.com/page3": 1,
				"https://example.com/page5": 1,
				"https://example.com/page1": 3,
			},
			expected : []pageCount{
				{
					normalizedURL: "https://example.com/page1",
					count: 3,
				},
				{
					normalizedURL: "https://example.com/page2",
					count: 2,
				},
				{
					normalizedURL: "https://example.com/page4",
					count: 2,
				},
				{
					normalizedURL: "https://example.com/page3",
					count: 1,
				},
				{
					normalizedURL: "https://example.com/page5",
					count: 1,
				},
			},
		},
		{
			name: "empty map",
			inputPages: map[string]int{},
			expected: []pageCount{},
		},
		{
			name: "nil map",
			inputPages: nil,
			expected: []pageCount{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.inputPages)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected pages: %v, actual: %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}