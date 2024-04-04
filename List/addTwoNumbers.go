package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	flag := false

	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		value := v1 + v2
		if flag {
			value++
			flag = false
		}
		if value >= 10 {
			flag = true
		}
		value %= 10

		nowNode := &ListNode{Val: value}
		tmp.Next = nowNode
		tmp = tmp.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag {
		last := &ListNode{Val: 1}
		tmp.Next = last
	}

	return head.Next
}
