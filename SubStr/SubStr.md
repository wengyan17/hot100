# SubStr
[TOC]

## 和为 K 的子数组
> Problem: [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/description/)

## 思路

> 可以进行n方暴力枚举区间，用前缀和算区间值，在此基础上进行优化，用哈希对前缀值进行桶计数，维护ans

## 解题方法

> 在n方的基础上
pre[i] - pre[j - 1] == k //这是前缀和的式子
将其进行移项得到：
pre[i] - k == pre[j - 1]
在这一步中，$$pre_i$$和$$pre_{j-1}$$都是前缀和，也就是说，只需要在遍历i的同时，维护$$pre_{j-1}$$的个数就行

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$

## Code
```Go []
func subarraySum(nums []int, k int) (ans int) {
	mp := map[int]int{}
	pre := 0
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if mp[pre-k] != 0 {
			ans += mp[pre-k]
		}
		mp[pre]++
	}
	return
}
```


##  滑动窗口最大值
> Problem: [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/description/)

## 思路

> 1.可以使用堆进行暴力筛
2.可以使用单调队列，维护单调队列每次输出队头

## 解题方法

> 这里附上官方题解链接
[官方题解](https://leetcode.cn/problems/sliding-window-maximum/solutions/543426/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/?envType=study-plan-v2&envId=top-100-liked)
这里主要学习golang中堆的使用方法

## 复杂度

时间复杂度:
> $O(nlog(n))$

空间复杂度:
> $O(n)$



## Code
```Go []
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}
```

## 最小覆盖子串
> Problem: [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/description/)


## 思路

> 滑动窗口维护子串字符，和字符串t进行比较，更新ans

## 解题方法

> 双指针模拟滑动窗口，左端点当不是t的字符时推出，判定通过时也推出，右端点则一直入队即可

## 复杂度

时间复杂度:
> $O(n)$

空间复杂度:
> $O(n)$



## Code
```Go []
func check(mps, mpt map[byte]int) bool {
	for k, v := range mpt {
		if mps[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) (ans string) {
	mps := map[byte]int{}
	mpt := map[byte]int{}

	for i := 0; i < len(t); i++ {
		mpt[t[i]]++
	}

	ansl, ansr := 0, math.MaxInt32

	for l, r := 0, 0; r < len(s); r++ {
		mps[s[r]]++

		for mpt[s[l]] == 0 {
			mps[s[l]]--
			if l < len(s)-1 {
				l++
			} else {
				break
			}
		}

		//for k, v := range mps {
		//	fmt.Println(k, " ", v)
		//}
		//fmt.Println(s[l:r+1], " ", check(mps, mpt))
		for check(mps, mpt) {
			if (r - l) < (ansr - ansl) {
				ansr, ansl = r, l
			}
			mps[s[l]]--
			l++
		}
	}

	if ansr == math.MaxInt32 {
		return ""
	}
	ans = s[ansl : ansr+1]
	return
}
```
  




