package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	value := []int{}
	for ; head != nil; head = head.Next {
		value = append(value, head.Val)
	}

	l, r := 0, len(value)-1
	for l < r {
		if value[l] != value[r] {
			return false
		}
		l++
		r--
	}

	return true
}
