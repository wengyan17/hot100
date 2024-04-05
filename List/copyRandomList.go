package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var mp = map[*Node]*Node{}

func copy(head *Node) *Node {
	if head == nil {
		return nil
	}

	if v, has := mp[head]; has {
		return v
	}
	newNode := &Node{head.Val, nil, nil}
	mp[head] = newNode
	newNode.Next = copy(head.Next)
	newNode.Random = copy(head.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	return copy(head)
}
