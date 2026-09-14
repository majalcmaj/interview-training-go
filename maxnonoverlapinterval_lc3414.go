package main

import (
	"sort"
)

const MINUS_INF = -1e18
const MAX_INTERVALS = 4

type interval struct {
	l, r, weight, idx int
}

func maximumWeight(rawIntervals [][]int) []int {
	intervals := make([]interval, len(rawIntervals))
	for idx, rawInterval := range rawIntervals {
		intervals[idx] = interval{rawInterval[0], rawInterval[1], rawInterval[2], idx}
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].r < intervals[j].r })

	dp := make([][]int, len(rawIntervals)+1)
	for i := 0; i <= len(rawIntervals); i++ {
		dp[i] = make([]int, MAX_INTERVALS+1)
	}
	for j := 1; j <= MAX_INTERVALS; j++ {
		dp[0][j] = MINUS_INF
	}

	biggestNonOverlapIdx := make([]int, len(rawIntervals))
	for i := 0; i < len(rawIntervals); i++ {
		biggestNonOverlapIdx[i] = sort.Search(i, func(idx int) bool { return intervals[i].l <= intervals[idx].r })
	}

	for i := 1; i <= len(rawIntervals); i++ {
		for j := 1; j <= MAX_INTERVALS; j++ {
			dp[i][j] = max(dp[biggestNonOverlapIdx[i-1]][j-1]+intervals[i-1].weight, dp[i-1][j])
		}
	}

	maxVal := -1
	bestJ := -1
	for j := 1; j <= MAX_INTERVALS; j++ {
		if dp[len(intervals)][j] > maxVal {
			bestJ = j
			maxVal = dp[len(intervals)][j]
		}
	}
	result := make([]int, 0, MAX_INTERVALS)

	for i, j := len(intervals), bestJ; j > 0 && i > 0; {
		if dp[i][j] == dp[i-1][j] {
			i = i - 1
		} else {
			result = append(result, intervals[i-1].idx)
			i = biggestNonOverlapIdx[i-1]
			j = j - 1
		}
	}
	sort.Ints(result)
	return result
}
