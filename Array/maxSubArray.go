package main

import "fmt"

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

func main() {
	nums := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
}
