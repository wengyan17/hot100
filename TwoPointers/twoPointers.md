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


## 三数之和
> Problem: [15. 三数之和](https://leetcode.cn/problems/3sum/description/)

## 思路

> 看了数据范围，判断平方级别，可以考虑双指针或者二分写法，这里用双指针

## 解题方法

> l，r钉住左右端点，当最大值+最小值小于目标值时，说明这个最小值不够大，中间都不要，l++，当最大值+最小值大于目标值时，说明这个最大值不够小，r--，外层遍历每个数，当做目标值即可

## 复杂度

时间复杂度:
> $O(n^2)$

空间复杂度:
> $O(n)$



## Code
```Go []
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println("nums:", nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[l] + nums[r]
			//fmt.Println("target:", target, "sum:", sum, "l, r:", l, r)
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				res := []int{nums[i], nums[l], nums[r]}
				//fmt.Println("res", res)
				ans = append(ans, res)
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return ans
}
```


## 接雨水
> Problem: [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/description/)


## 思路

> 把金字塔部分求出来，然后减去黑方块部分

## 解题方法

> 双指针钉住左右端点，更新左右的最大值，按照短板来计算桶能装水的数量，哪边短哪边就加入ans，然后端点移动

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(1)$



## Code
```Go []
func trap(height []int) int {
	l := 0
	r := len(height) - 1

	maxpre := 0
	maxsuf := 0
	ans := 0

	for l <= r {
		maxpre = max(maxpre, height[l])
		maxsuf = max(maxsuf, height[r])

		if maxpre <= maxsuf {
			ans += min(maxpre, maxsuf) - height[l]
			l++
		} else {
			ans += min(maxpre, maxsuf) - height[r]
			r--
		}
	}

	return ans
}
```







