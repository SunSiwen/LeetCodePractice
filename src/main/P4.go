package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		var tmp = nums1
		nums1 = nums2
		nums2 = tmp
	}

	m, n := len(nums1), len(nums2)
	total := m + (n-m+1)>>1
	l, r := 0, m
	for l < r {
		mid := l + (r-l+1)>>1
		rest := total - mid
		if nums1[mid-1] > nums2[rest] {
			r = mid - 1
		} else {
			l = mid
		}
	}
	var l1, l2, r1, r2 int
	r = total - l
	if l == 0 {
		l1 = math.MinInt64
	} else {
		l1 = nums1[l-1]
	}
	if l == m {
		r1 = math.MaxInt64
	} else {
		r1 = nums1[l]
	}
	if r == 0 {
		l2 = math.MinInt64
	} else {
		l2 = nums2[r-1]
	}
	if r == n {
		r2 = math.MaxInt64
	} else {
		r2 = nums2[r]
	}
	if (m+n)%2 == 1 {
		return float64(maxInt(l1, l2))
	} else {
		return float64(maxInt(l1, l2)+minInt(r1, r2)) / 2
	}
}

func minInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

//func maxInt(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
