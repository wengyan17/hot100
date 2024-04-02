package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	mp := map[*ListNode]bool{}

	for nowNode := head; nowNode != nil; nowNode = nowNode.Next {
		if mp[nowNode] {
			return nowNode
		}
		mp[nowNode] = true
	}

	return nil
}
