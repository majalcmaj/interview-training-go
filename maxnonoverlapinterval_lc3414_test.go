package main

import (
	"slices"
	"testing"
)

func TestMaximumWeight(t *testing.T) {
	for _, d := range []struct {
		intervals [][]int
		expected  []int
	}{
		{[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}, {6, 9, 3}, {6, 7, 1}, {8, 9, 1}}, []int{2, 3}},
		{[][]int{{5, 8, 1}, {6, 7, 7}, {4, 7, 3}, {9, 10, 6}, {7, 8, 2}, {11, 14, 3}, {3, 5, 5}}, []int{1, 3, 5, 6}},
		{[][]int{{1, 2, 20}}, []int{0}},
		{[][]int{{1, 100, 100}, {1, 40, 50}, {41, 100, 49}}, []int{0}},
		{[][]int{{1, 100, 100}, {1, 40, 50}, {41, 100, 51}}, []int{1, 2}},
	} {
		result := maximumWeight(d.intervals)
		if !slices.Equal(result, d.expected) {
			t.Errorf(`Expected %v got %v for input %v`, d.expected, result, d.intervals)
		}
	}
}
