package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, curr *ListNode
	carr := 0
	for l1 != nil || l2 != nil || carr != 0 {
		mod := 0
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + carr
		mod, carr = sum%10, sum/10
		if res == nil {
			res = &ListNode{Val: mod}
			curr = res
		} else {
			curr.Next = &ListNode{Val: mod}
			curr = curr.Next
		}
	}
	return res
}
