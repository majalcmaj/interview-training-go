package main

const MOD_BY = 1e9 + 7

func _numberOfSets(n int, k int) int64 {
	dp := make([][]int64, k+1)

	for i := 0; i < k+1; i++ {
		dp[i] = make([]int64, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}

	pre := make([]int64, n+1)
	for i := 1; i <= k; i++ {
		pre[0] = dp[i-1][0]
		for t := 1; t <= n; t++ {
			pre[t] = (pre[t-1] + dp[i-1][t]) % MOD_BY
		}

		for j := i; j <= n; j++ {
			sum := (pre[j-1] - pre[i-1] + MOD_BY) % MOD_BY
			dp[i][j] = (dp[i][j-1] + sum) % MOD_BY
		}
	}

	return dp[k][n]
}

func numberOfSets(n int, k int) int {
	return int(_numberOfSets(n, k))
}
