package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode = nil
	nowNode := head
	for nowNode != nil {
		nextNode := nowNode.Next
		nowNode.Next = preNode
		preNode = nowNode
		nowNode = nextNode
	}
	return preNode
}
