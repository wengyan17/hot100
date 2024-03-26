# 滑动窗口

[TOC]

## 无重复字符的最长子串
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

## 找到字符串中所有字母异位词
> Problem: [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/)


## 思路

> 滑动窗口桶排计数

## 解题方法

> 滑动窗口维护一个长度为字符串p长度的窗口，对窗口内字符进行桶排，比较在窗口滑动时和p字符串字符出现次数是否相同，相同就把下标加入到ans返回，不同就跳过

## 复杂度

时间复杂度:
>  $O(n)$

空间复杂度:
> $O(26)$



## Code
```Go []
func findAnagrams(s string, p string) (ans []int) {
	ss, sp := len(s), len(p)
	if sp > ss {
		return
	}
	count := [26]int{}

	for i := 0; i < sp; i++ {
		count[s[i]-'a']++
		count[p[i]-'a']--
	}

	diff := 0
	for _, i := range count {
		if i != 0 {
			diff++
		}
	}
	if diff == 0 {
		ans = append(ans, 0)
	}

	for i := sp; i < ss; i++ {
		//fmt.Println(count)
		count[s[i]-'a']++
		if count[s[i]-'a'] == 0 {
			diff--
		} else if count[s[i]-'a'] == 1 {
			diff++
		}
		//fmt.Println(diff)

		count[s[i-sp]-'a']--
		if count[s[i-sp]-'a'] == 0 {
			diff--
		} else if count[s[i-sp]-'a'] == -1 {
			diff++
		}
		//fmt.Println(diff)

		if diff == 0 {
			ans = append(ans, i-sp+1)
		}
	}
	return
}
```
  


