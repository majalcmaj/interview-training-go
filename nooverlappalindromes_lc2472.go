package main

import "fmt"

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
	if k%2 == 0 {
		even := evenPalindromes(s, k)
		odd := oddPalindromes(s, k+1)

		fmt.Printf("O:%v\nE:%v\n", odd, even)

		lastFoundIdx := -k - 1
		count := 0
		for i := 0; i < len(even); i++ { // Is this number good?
			if (odd[i] || even[i]) && lastFoundIdx < i-k {
				lastFoundIdx = i
				count += 1
				odd[i] = true  // TODO - rm
				even[i] = true // TODO - rm
			} else {
				odd[i] = false  // TODO - rm
				even[i] = false // TODO - rm
			}
		}
		if odd[len(odd)-1] {
			count++
		}
		// fmt.Printf("O: %v\n\n %v\n", odd, even)
		return count
	} else {
		odd := oddPalindromes(s, k)
		even := evenPalindromes(s, k+1)

		fmt.Printf("O:%v\nE:%v\n", odd, even)
		lastFoundIdx := -k - 1
		count := 0
		for i := 0; i < len(even); i++ { // Is this number good?
			if (odd[i] || even[i]) && lastFoundIdx <= i-k {
				lastFoundIdx = i
				count += 1
				odd[i] = true // TODO - rm
			} else {
				odd[i] = false // TODO - rm
			}
		}
		if odd[len(odd)-1] {
			count++
		}
		return count
	}
}
