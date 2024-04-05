package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func checkKGroup(head *ListNode, k int) bool {
	for i := 0; i < k; i++ {
		if head.Next == nil {
			return false
		}
		head = head.Next
	}
	return true
}

func reverseGroup(head *ListNode, k int) *ListNode {
	var stack []*ListNode
	h := head
	tmp := h
	for i := 0; i <= k; i++ {
		stack = append(stack, head.Next)
		head = head.Next
	}

	for i := k - 1; i >= 0; i-- {
		tmp.Next = stack[i]
		tmp = tmp.Next
	}

	tmp.Next = stack[k]
	return h
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	h := &ListNode{0, head}
	tmp := h
	for checkKGroup(tmp, k) {
		reverseGroup(tmp, k)

		for i := 0; i < k; i++ {
			tmp = tmp.Next
		}
	}

	return h.Next
}
