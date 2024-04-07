package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap struct {
	NodeList []*ListNode
}

// Push 向堆中插入元素
func (h *MinHeap) Push(node *ListNode) {
	h.NodeList = append(h.NodeList, node)
	h.up(len(h.NodeList) - 1)
}

// Pop 从堆顶删除元素
func (h *MinHeap) Pop() *ListNode {
	peekedNode := h.NodeList[0]
	h.swap(0, len(h.NodeList)-1)
	h.NodeList = h.NodeList[:len(h.NodeList)-1]
	h.down(0)
	return peekedNode
}

// up 上浮
func (h *MinHeap) up(idx int) {
	for h.NodeList[idx].Val < h.NodeList[(idx-1)/2].Val {
		h.swap((idx-1)/2, idx)
		idx = (idx - 1) / 2
	}
}

// down 下沉
func (h *MinHeap) down(idx int) {
	for 2*idx+1 < len(h.NodeList) {
		left, right := 2*idx+1, 2*idx+2
		smallest := left
		if right < len(h.NodeList) && h.NodeList[right].Val < h.NodeList[left].Val {
			smallest = right
		}
		if h.NodeList[smallest].Val > h.NodeList[idx].Val {
			break
		}
		h.swap(smallest, idx)
		idx = smallest
	}
}

// swap 交换两个节点
func (h *MinHeap) swap(i, j int) {
	h.NodeList[i], h.NodeList[j] = h.NodeList[j], h.NodeList[i]
}

func sortList(head *ListNode) *ListNode {
	heap := MinHeap{}
	for ; head != nil; head = head.Next {
		heap.Push(head)
	}

	dummyHead := &ListNode{}
	currentNode := dummyHead
	for len(heap.NodeList) > 0 {
		topNode := heap.Pop()
		topNode.Next = nil
		currentNode.Next = topNode
		currentNode = currentNode.Next
	}

	return dummyHead.Next
}
