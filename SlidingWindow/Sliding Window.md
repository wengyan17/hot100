# 滑动窗口

[TOC]

> Problem: [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/)



## 思路

> 队列记录当前子串，用map记录队列里有哪些子串，遍历更新ans长度

## 解题方法

> 队列记录当前子串，用map记录队列里有哪些子串，遍历更新ans长度

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
>  $O(n)$



## Code
```Go []
func lengthOfLongestSubstring(s string) int {
	ans := 0
	var mp = map[uint8]bool{}
	var str = ""
	for i := 0; i < len(s); i++ {
		if mp[s[i]] {
			for mp[s[i]] {
				mp[str[0]] = false
				str = str[1:]
			}
			str = str + string(s[i])
			mp[s[i]] = true
			ans = max(ans, len(str))
		} else {
			str = str + string(s[i])
			mp[s[i]] = true
			ans = max(ans, len(str))
		}
	}
	return ans
}
```
  
