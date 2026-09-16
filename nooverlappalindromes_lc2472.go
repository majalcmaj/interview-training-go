package main

func oddPalindromes(s string, k int) []bool {
	margin := k / 2
	dp := make([]bool, len(s))
	if len(s) < k {
		return dp
	}

	for i := margin; i < len(dp)-margin; i++ {
		dp[i] = true
	}

	for i := 1; i <= margin; i += 1 { // is this correct? Should be 1?
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
	if len(s) < k {
		return dp
	}

	if k == 1 { // TODO: needed?
		return dp
	}

	for i := margin - 1; i < len(dp); i++ {
		dp[i] = true
	}

	for i := 0; i < margin; i++ {
		for j := i; j < len(s)-i-1; j++ {
			if !(dp[j] && s[j-i] == s[j+i+1]) {
				dp[j] = false
			}
		}
	}
	return dp
}

func maxPalindromes(s string, k int) int {
	n := len(s)

	// Smallest odd length >= k, and smallest even length >= k.
	evenK := k + k%2
	odd := oddPalindromes(s, k)
	even := evenPalindromes(s, evenK)

	marginOdd := k / 2
	marginEven := evenK / 2
	lenOdd := 2*marginOdd + 1
	lenEven := 2 * marginEven

	lastEnd := -1
	count := 0
	for end := 0; end < n; end++ {
		i := end - marginOdd
		oddOK := i >= 0 && i < len(odd) && odd[i] && end-lenOdd+1 > lastEnd

		j := end - marginEven
		evenOK := j >= 0 && j < len(even) && even[j] && end-lenEven+1 > lastEnd

		if oddOK || evenOK {
			count++
			lastEnd = end
		}
	}
	return count
}
