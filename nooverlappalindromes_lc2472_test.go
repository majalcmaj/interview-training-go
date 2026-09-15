package main

import "testing"

func TestMaxOverlappingPalindromes(t *testing.T) {
	for _, d := range []struct {
		s           string
		k, expected int
	}{
		{"abaccdbbd", 3, 2},
		{"adbcda", 2, 0},
		{"abaaba", 3, 2},
		{"aaaaaaaaa", 3, 3},
		{"abaccdbbd", 1, 9},
	} {
		res := maxPalindromes(d.s, d.k)
		if res != d.expected {
			t.Errorf("Expected %d, got %d for s=%s, k=%d", d.expected, res, d.s, d.k)
		}
	}
}
