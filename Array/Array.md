# 普通数组
[TOC]

## 最大子数组和
> Problem: [53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/description/)

## 思路

> 用f[i]表示以i结尾的最大连续子数组，O(n)递归

## 解题方法

> f[i] = max(f[i - 1] + nums[i], nums[i])
注：当以i结尾时，nums[i]是已经要选的，即使是负的

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$


## Code
```Go []
func maxSubArray(nums []int) (ans int) {
	var n = len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	ans = f[0]
	for i := 1; i < n; i++ {
		f[i] = max(f[i-1]+nums[i], nums[i])
		ans = max(ans, f[i])
	}
	return ans
}
```

## 合并区间
> Problem: [56. 合并区间](https://leetcode.cn/problems/merge-intervals/description/)

## 思路

> 按照区间左值排序，然后遍历后续区间有无重叠，重叠就合并，没重叠就是新区间

## 解题方法

> 按照区间左值排序，然后遍历后续区间有无重叠，重叠就合并，没重叠就是新区间

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	n := len(intervals)
	for i := 1; i < n; i++ {
		last := len(ans) - 1
		if intervals[i][0] <= ans[last][1] {
			ans[last][1] = max(ans[last][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return
}
```


## 轮转数组
> Problem: [189. 轮转数组](https://leetcode.cn/problems/rotate-array/description/)

## 思路

> 切下数组的后k位放到数组前面

## 解题方法

> 使用golang中的切片，然后append

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func rotate(nums []int, k int) {
	k = k % len(nums)
	newNums := append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	copy(nums, newNums)
}
```


## 除自身以外数组的乘积
> Problem: [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/description/)

## 思路

> 算出总乘积，除以当前值即可，但是注意有0需要特判

## 解题方法

> 分类讨论0的数量去特判就行

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func productExceptSelf(nums []int) (ans []int) {
	sum := 1
	numsZore := 0
	for _, i := range nums {
		if i == 0 {
			numsZore++
		} else {
			sum *= i
		}
	}

	if numsZore > 1 {
		for i := 0; i < len(nums); i++ {
			ans = append(ans, 0)
		}
	} else if numsZore == 1 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				ans = append(ans, sum)
			} else {
				ans = append(ans, 0)
			}
		}
	} else if numsZore == 0 {
		for _, i := range nums {
			ans = append(ans, sum/i)
		}
	}

	return
}
```

## 缺失的第一个正数
> Problem: [41. 缺失的第一个正数](https://leetcode.cn/problems/first-missing-positive/description/)


## 思路
> 裸的寻找mex值

## 解题方法
> 桶排，然后遍历，小于0的不用统计，大于n的也不用统计，因为如果有n个数，但是有大于n的值，那么mex值一定小于n

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func firstMissingPositive(nums []int) int {
	cnt := make([]int, len(nums)+2)

	for _, v := range nums {
		if v < len(nums)+2 && v >= 0 {
			cnt[v]++
		}
	}

	ans := 1
	for i := 1; i < len(nums)+2; i++ {
		if cnt[i] == 0 {
			ans = i
			break
		}
	}

	return ans
}
```



