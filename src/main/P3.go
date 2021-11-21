package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res, l, r := 1, 0, 1
	for r < len(s) {
		tmp := s[l:r]
		if -1 == strings.Index(tmp, string(s[r])) {
			res = maxInt(res, r-l+1)
		} else {
			index := strings.Index(tmp, string(s[r]))
			l = l + index + 1
		}
		r++
	}
	res = maxInt(r-l, res)
	return res
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
