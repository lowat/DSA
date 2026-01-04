package main

import "testing"

func TestValidAnagram(t *testing.T) {
	tests := []struct {
		s string
		t string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "a", true},
		{"a", "b", false},
		{"a", "aa", false},
		{"abc", "ab", false},
		{"aabb", "bbaa", true},
		{"aabb", "abbb", false},
		{"listen", "silent", true},
	}

	for _, test := range tests {
		result := IsAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
