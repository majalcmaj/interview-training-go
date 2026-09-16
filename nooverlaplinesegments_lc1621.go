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

	for i := 1; i <= k; i++ {
		for j := i; j <= n; j++ {
			dp[i][j] = dp[i][j-1]
			for l := i; l < j; l++ {
				dp[i][j] = (dp[i][j] + dp[i-1][l]) % MOD_BY
			}
		}
	}

	return dp[k][n]
}

func numberOfSets(n int, k int) int {
	return int(_numberOfSets(n, k))
}
