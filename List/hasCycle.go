package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	mp := map[*ListNode]bool{}

	for nowNode := head; nowNode != nil; nowNode = nowNode.Next {
		if mp[nowNode] {
			return true
		}
		mp[nowNode] = true
	}

	return false
}
