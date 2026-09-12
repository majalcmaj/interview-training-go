package main

type ctx [][]int

func (c ctx) maxIntervalsRecursive(idx int, curr []int) ([]int, int) {
	idxs := make([]int, len(curr)+1)
	copy(idxs, curr)
	idxs[len(curr)] = idx

	bestIdxs := idxs
	bestTotal := c.sumWeights(idxs)

	if len(idxs) == 4 || idx >= len(c) {
		return bestIdxs, bestTotal
	}

	for i := idx + 1; i < len(c); i++ {
		if c.isNoOverlap(i, idxs) {
			newBestIdxs, newBestTotal := c.maxIntervalsRecursive(i, idxs)
			if newBestTotal > bestTotal {
				bestIdxs = newBestIdxs
				bestTotal = newBestTotal
			}
		}
	}
	return bestIdxs, bestTotal
}

func (c ctx) sumWeights(idxs []int) int {
	sum := 0
	for _, idx := range idxs {
		sum += c[idx][2]
	}
	return sum
}

func (c ctx) isNoOverlap(idx int, curr []int) bool {
	for _, curIdx := range curr {
		if !(c[idx][0] > c[curIdx][1] || c[curIdx][0] > c[idx][1]) {
			return false
		}
	}
	return true
}

func maximumWeight(intervals [][]int) []int {
	c := ctx(intervals)
	idxs := make([]int, 0)
	bestTotal := -1
	result := []int{}
	for i := 0; i < len(intervals); i++ {
		newRes, newTotal := c.maxIntervalsRecursive(i, idxs)
		if newTotal > bestTotal {
			result = newRes
			bestTotal = newTotal
		}
	}
	return result
}
