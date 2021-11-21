package main

//func longestPalindrome(s string) string {
//	i := len(s)
//	if i < 2 {
//		return s
//	}
//	res := s[0:1]
//	dp := make([][]bool, i)
//	for index := range dp {
//		dp[index] = make([]bool, i)
//	}
//	for start := 0; start < i; start++ {
//		dp[start][start] = true
//	}
//
//	for start := len(dp) - 1; start >= 0; start-- {
//		for end := start + 1; end < len(dp); end++ {
//			if end == start+1 {
//				dp[start][end] = s[start] == s[end]
//			} else {
//				dp[start][end] = (s[start] == s[end]) && dp[start+1][end-1]
//			}
//
//			if end-start+1 > len(res) {
//				res = s[start : end+1]
//			}
//		}
//	}
//	return res
//}

func longestPalindrome(s string) string {
	i := len(s)
	if i < 2 {
		return s
	}
	res := s[0:1]
	dp := make([]bool, i)

	dp[i-1] = true

	for start := len(dp) - 2; start >= 0; start-- {
		for end := len(dp) - 1; end >= start; end-- {
			if end <= start+1 {
				dp[end] = s[start] == s[end]
			} else {
				dp[end] = (s[start] == s[end]) && dp[end-1]
			}

			if dp[end] && (end-start+1 > len(res)) {
				res = s[start : end+1]
			}
		}
	}
	return res
}
