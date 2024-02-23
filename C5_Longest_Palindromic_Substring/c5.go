package c5

import "strings"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	begin, end := 0, 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	dp[0][0] = true
	for right := 1; right < len(s); right++ {
		dp[right][right] = true
		for left := 0; left < right; left++ {
			if s[left] == s[right] && (right-left <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true
				if right-left > end-begin {
					begin = left
					end = right
				}
			}
		}
	}
	return s[begin : end+1]
}

func longestPalindromeManacher(s string) string {
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)
	P := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		if R > i {
			P[i] = min(R-i, P[2*C-i])
		}
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}

	maxLen := 0
	centerIndex := 0
	for i, v := range P {
		if v > maxLen {
			maxLen = v
			centerIndex = i
		}
	}
	return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}
