package main

func oddPalindromes(s string, k int) []bool {
	margin := k / 2
	dp := make([]bool, len(s))

	for i := margin; i < len(dp)-margin; i++ {
		dp[i] = true
	}

	for i := 1; i <= margin; i += 2 { // is this correct? Should be 1?
		for j := i; j < len(s)-i; j++ {
			if !(dp[j] && s[j-i] == s[j+i]) {
				dp[j] = false
			}
		}
	}
	return dp
}

func evenPalindromes(s string, k int) []bool {
	margin := k / 2
	dp := make([]bool, len(s)-1)

	if k == 1 { // TODO: needed?
		return dp
	}

	for i := margin - 1; i < len(dp); i++ {
		dp[i] = true
	}

	for i := 0; i <= margin; i++ {
		for j := i; j < len(s)-i-1; j++ {
			if !(dp[j] && s[j-i] == s[j+i+1]) {
				dp[j] = false
			}
		}
	}
	return dp
}

func maxPalindromes(s string, k int) int {

	return 1
}
