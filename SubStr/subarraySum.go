package main

import "fmt"

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

func main() {
	nums := []int{1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}
