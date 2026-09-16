package main

func _numberOfSets(n int, k int) int64 {
	// fmt.Printf("n %d k %d\n", n, k)
	if k == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	sum := _numberOfSets(n-1, k)
	for i := k; i < n; i++ {
		sum += _numberOfSets(i, k-1)
	}
	return sum
}

func numberOfSets(n int, k int) int {
	return int(_numberOfSets(n, k) % (1e9 + 7))
}
