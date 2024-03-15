# hash
## 两数之和
>  Problem: [1. 两数之和](https://leetcode.cn/problems/two-sum/description/)

[TOC]

## 思路

> 桶排数组，使用map存储

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func twoSum(nums []int, target int) []int {
	mp := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = append(mp[nums[i]], i)
	}

	var ans []int
	flag := false
	for i := 0; i < len(nums); i++ {
		if flag {
			break
		}
		fst := nums[i]
		for _, num := range mp[target-fst] {
			if num != i {
				flag = true
				ans = append(ans, i, num)
				break
			}
		}
	}
	sort.Ints(ans)
	return ans
}
```


## 字母异位词分组
> ## Problem: [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/description/)


## 思路

> 哈希桶排

## 解题方法

> 利用切片string排序，整合成有序装进map，最后转化成二维数组进行输出
## 复杂度

时间复杂度:
> $O(strs.length * strs[i].length)$

空间复杂度:
> $O(n)$



## Code
```Go []
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, i := range strs {
		sli := strings.Split(i, "")
		sort.Strings(sli)
		sortcli := strings.Join(sli, "")
		mp[sortcli] = append(mp[sortcli], i)
	}

	var ans [][]string
	for _, i := range mp {
		ans = append(ans, i)
	}
	return ans
}
```

# 最长连续序列
> Problem: [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/description/)

## 思路

> 原本是想桶排然后依次排查，但是这样时间复杂度超过o(N)，优化一下桶排之后对数字的存在进行讨论，优化几次之后通过

## 解题方法

> 桶排数组，然后遍历数字，如果该数字前面还有头，则continue，如果该数字不是最后一个出现的，则continue，这样保证数组里的每个数字都有且只遍历一遍

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func longestConsecutive(nums []int) int {
	var ans = 0
	var mp = map[int]int{}
	for _, i := range nums {
		mp[i]++
	}
	for _, i := range nums {
		if mp[i-1] > 0 {
			continue
		}
		if mp[i] > 1 {
			mp[i]--
			continue
		}

		j := i + 1
		for {
			if mp[j] > 0 {
				j++
			} else {
				break
			}
		}
		ans = max(ans, j-i)
	}

	return ans
}
```
  


