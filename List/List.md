# 链表

[TOC]

## 相交链表
> Problem: [160. 相交链表](https://leetcode.cn/problems/intersection-of-two-linked-lists/description/)



## 思路
>  使用桶记录A有哪些节点，遍历B寻找第一个相同的

## 解题方法
> 使用桶记录A有哪些节点，遍历B寻找第一个相同的并返回

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$

## Code
```Go []
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
```

## 反转链表
> Problem: [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/description/)


## 思路

> 记录前一节点，当遍历至后一节点时将该节点的next指向前一节点

## 解题方法

> 像swap那样交换就行

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(1)$



## Code
```Go []
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
```

## 回文链表
> Problem: [234. 回文链表](https://leetcode.cn/problems/palindrome-linked-list/description/)

## 思路

> 取出值，然后判断数组是否是回文即可

## 解题方法

> 取出值，然后判断数组是否是回文即可

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
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
		l ++
		r --
	}

	return true
}
```


## 环形链表
> Problem: [141. 环形链表](https://leetcode.cn/problems/linked-list-cycle/description/)


## 思路

> 将节点加入map判重

## 解题方法

> 将节点加入map判重

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func hasCycle(head *ListNode) bool {
	mp := map[*ListNode]bool{}
	
	for nowNode := head;nowNode!= nil; nowNode = nowNode.Next {
		if mp[nowNode] {
			return true
		}
		mp[nowNode] = true
	}
	
	return false
}
```

## 环形链表 II
> Problem: [142. 环形链表 II](https://leetcode.cn/problems/linked-list-cycle-ii/description/)


## 思路

> 用map记录节点查重，遇到重复的就返回

## 解题方法

> 用map记录节点查重，遇到重复的就返回

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
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
```


## 合并两个有序链表
> Problem: [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/description/)

## 思路

> 是一个模拟，最初想法是用map都给存起来，map[val]*Node这样，但是感觉麻烦了

## 解题方法

> 开一条新的链，哪边小就先接上去，这样可以保证有序

## 复杂度

时间复杂度:
> $O(min(n, m))$

空间复杂度:
> $O(n)$



## Code
```Go []
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}

	tmp := head

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			tmp.Next = list2
			list2 = list2.Next
		} else {
			tmp.Next = list1
			list1 = list1.Next
		}
		tmp = tmp.Next
	}

	if list1 != nil {
		tmp.Next = list1
	}
	if list2 != nil {
		tmp.Next = list2
	}

	return head.Next
}
```

## 两数相加
> Problem: [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/description/)


## 思路

> 模拟

## 解题方法

> 使用一条新的链来装结果，注意“进位”的模拟即可

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
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
```


## 删除链表的倒数第 N 个结点
> Problem: [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/)

## 思路

> 算长度，删除倒数第n个点

## 解题方法

> 遍历求长度然后遍历删除点

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(1)$



## Code
```Go []
func getListLen(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getListLen(head)
	dummy := &ListNode{0, head}
	tmp := dummy
	for i := 0; i < length-n; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return dummy.Next
}
```

## 两两交换链表中的节点
> Problem: [24. 两两交换链表中的节点](https://leetcode.cn/problems/swap-nodes-in-pairs/description/)

## 思路

> 按照题目意思模拟

## 解题方法

> 注意两点交换越界的细节即可

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(1)$



## Code
```Go []
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
```


## K 个一组翻转链表
> Problem: [25. K 个一组翻转链表](https://leetcode.cn/problems/reverse-nodes-in-k-group/description/)


## 思路

> 大模拟，按照题意进行模拟即可

## 解题方法

> 细节方面
我将这题几个板块进行了打包封装成函数
1.有k个的进行调换，没有k个的不变---这一步就需要判断
我封装成了checkKGroup函数
检查接下来k个是不是nil
2.调换逻辑
我封装成了reverseGroup
将这k个进行翻转，然后再接上去

## 复杂度

时间复杂度:
> 添加时间复杂度, 示例： $O(n)$

空间复杂度:
> 添加空间复杂度, 示例： $O(n)$



## Code
```Go []
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
	for i := 0; i <= k; i ++ {
		stack = append(stack, head.Next)
		head = head.Next
	}
	
	for i := k - 1; i >= 0; i -- {
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
```

##  随机链表的复制
> Problem: [138. 随机链表的复制](https://leetcode.cn/problems/copy-list-with-random-pointer/description/)

## 思路

> 就如题意所说，他需要的是一个深拷贝，初步看见题目需要捋一下这题的难点在哪
1.Node内设了一个random值，当这个节点出来时，random节点可能还没遍历到，不好处理，也就是处理后面的节点的时候需要“回头”
2.他需要一条全新的链，而不是直接返回原来的head节点

## 解题方法

> 因为需要回头，所以递归就是非常容易考虑到的顺向思维，再使用map来存储已经出现过的节点即可

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
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
```

## 排序链表
> Problem: [148. 排序链表](https://leetcode.cn/problems/sort-list/description/)

## 思路

> 简单的链表排序,可以使用普通的冒泡,这里为了熟悉goalng堆的应用选择了堆排

## 解题方法

> 手写堆,然后节点进堆再出堆

## 复杂度

时间复杂度:
> $O(nlogn)$

空间复杂度:
> $O(n)$



## Code
```Go []
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
```
  






