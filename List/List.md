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





