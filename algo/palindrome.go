package algo

import (
	"regexp"
)

func constrainString(s string) string {
	// Define a regular expression pattern allowing only digits and English letters
	re := regexp.MustCompile("[^a-zA-Z0-9]")

	// Use the regular expression to replace any non-digit and non-letter characters with an empty string
	constrained := re.ReplaceAllString(s, "")

	return constrained
}

func LongestPalindrome(s string) string {
	if len(s) < 1 || len(s) > 1000 {
		return s
	}
	s = constrainString(s)
	start, maxLen := 0, 1
	n := len(s)

	// Create a 2D slice to store the palindrome information
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// All substrings of length 1 are palindromes
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Check substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// Check substrings of length 3 or more
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLen = length
			}
		}
	}

	return s[start : start+maxLen]
}
