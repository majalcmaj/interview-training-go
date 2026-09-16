package main

import "testing"

func TestNumberOfSets(t *testing.T) {
	for _, d := range []struct {
		n, k, expect int
	}{
		{4, 2, 5},
		{3, 1, 3},
		{30, 7, 796297179},
	} {
		res := numberOfSets(d.n, d.k)
		if res != d.expect {
			t.Errorf("Expected %d, got %d for n=%d, k=%d", d.expect, res, d.n, d.k)
		}
	}
}
