package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{0, head}
	tmp := h
	for tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := tmp.Next.Next
		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		tmp = node1
	}

	return h.Next
}
