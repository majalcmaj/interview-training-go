package main

import "testing"

func TestMinSumOfLengths(t *testing.T) {
	for _, d := range []struct {
		arr    []int
		target int
		expect int
	}{
		{[]int{3, 2, 2, 4, 3}, 3, 2},
		{[]int{7, 3, 4, 7}, 7, 2},
		{[]int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
		{[]int{3, 2, 2, 4, 1, 3}, 5, 4},
		{[]int{1, 1, 1, 1, 1, 1}, 3, 6},
	} {
		res := minSumOfLengths(d.arr, d.target)
		if res != d.expect {
			t.Errorf("Expected %d, got %d for arr=%v, target=%d", d.expect, res, d.arr, d.target)
		}
	}
}
