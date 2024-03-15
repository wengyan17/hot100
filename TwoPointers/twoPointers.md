# TwoPointers

[TOC]

## 移动零
> Problem: [283. 移动零](https://leetcode.cn/problems/move-zeroes/description/)


## 思路

> 双指针swap

## 解题方法

> 双指针swap，将后面的非0数一个个按顺序换到前面来

## 复杂度

时间复杂度:
>  $O(n)$

空间复杂度:
>  $O(n)$



## Code
```Go []
func moveZeroes(nums []int) {
	noZ := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[noZ], nums[r] = nums[r], nums[noZ]
			//fmt.Println(nums)
			noZ++
		}
	}
}
```

## 盛最多水的容器
> Problem: [11. 盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/)


## 思路

> 先考虑暴力做法，也就是N方的遍历思路，然后优化其中重复计算的部分，也就是两柱子之间小于等于边界柱子的柱子一定不能是答案的那两个，所以双指针遍历边界即可

## 解题方法

> 双指针遍历边界，哪边小移动哪边，更新答案输出

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0

	for l < r {
		res := min(height[l], height[r]) * (r - l)
		//fmt.Println(res)
		ans = max(ans, res)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}
```



