package main

import (
	"testing"
)

func TestFindLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abcabcbb", "abc"},
		{"bbbb", "b"},
		{"pwwkew", "wke"},
		{"", ""},
		{"abcdefg", "abcdefg"},
		{"a", "a"},
	}

	for _, test := range tests {
		result := findSubstring(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
