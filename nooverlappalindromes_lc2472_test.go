package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestOddPalindromes(t *testing.T) {
	for _, d := range []struct {
		s      string
		k      int
		expect []bool
	}{
		{"gataghmwwmoeyeov", 5, genExpected(16, 2, 12)},
		{"abaccdbbd", 3, []bool{false, true, false, false, false, false, false, false, false}},
		{"adbcda", 3, []bool{false, false, false, false, false, false}},
		{"ababa", 5, []bool{false, false, true, false, false}},
		{"abaca", 5, []bool{false, false, false, false, false}},
		{"aaaaaa", 3, []bool{false, true, true, true, true, false}},
	} {
		res := oddPalindromes(d.s, d.k)
		if !slices.Equal(res, d.expect) {
			t.Errorf("Expected %v, got %v for s=%s, k=%d", d.expect, res, d.s, d.k)
		}
	}
}

func TestEvenPalindromes(t *testing.T) {
	for idx, d := range []struct {
		s      string
		k      int
		expect []bool
	}{
		{"gataghmwwmoeyeov", 4, genExpected(15, 7)},
		{"abaccdbbd", 4, []bool{false, false, false, false, false, false, true, false}},
		{"adbcda", 2, []bool{false, false, false, false, false}},
		{"abaaba", 6, []bool{false, false, true, false, false}},
		{"aaaaaa", 2, []bool{true, true, true, true, true}},
	} {
		fmt.Printf("Running %d\n", idx)
		res := evenPalindromes(d.s, d.k)
		if !slices.Equal(res, d.expect) {
			t.Errorf("Expected %v, got %v for s=%s, k=%d", d.expect, res, d.s, d.k)
		}
	}
}

func TestMaxOverlappingPalindromes(t *testing.T) {
	for _, d := range []struct {
		s           string
		k, expected int
	}{
		{"qvlbphaahpblshrwzzwrhscaivaeffllffeanwtlaaltwngtupwkwwkwpuaa", 10, 5},
		{"aa", 2, 1},
		{"baa", 3, 0},
		{"fttfjofpnpfydwdwdnns", 2, 4},
		{"gataghmwwmoeyeov", 4, 3},
		{"abaccdbbd", 3, 2},
		{"abaaba", 3, 2},
		{"abaccdbbd", 1, 9},
		{"adbcda", 2, 0},
		{"aaaaaaaaa", 3, 3},
		{"cabaccabac", 4, 2},
		{"aaaaaaaaa", 4, 2},
	} {
		res := maxPalindromes(d.s, d.k)
		if res != d.expected {
			t.Errorf("Expected %d, got %d for s=%s, k=%d", d.expected, res, d.s, d.k)
		}
	}
}

func genExpected(l int, trues ...int) []bool {
	res := make([]bool, l)
	for _, idx := range trues {
		res[idx] = true
	}
	return res
}
