package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	flag := map[*ListNode]bool{}

	for node := headA; node != nil; node = node.Next {
		flag[node] = true
	}

	for node := headB; node != nil; node = node.Next {
		if flag[node] {
			return node
		}
	}
	return nil
}
