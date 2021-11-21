package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		a, ok := m[v]
		if ok {
			res:= []int{i, a}
			return res
		} else {
			m[target-v] = i
		}
	}
	res := []int{-1,-1}
	return res
}