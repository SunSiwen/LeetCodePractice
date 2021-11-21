package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("hello")

	twoSum([]int{2, 7, 11, 15}, 9)
	addTwoNumbers(nil, nil)
	lengthOfLongestSubstring("abcabcbb")
	findMedianSortedArrays([]int{1, 2}, []int{0})
	longestPalindrome("abc")
}
