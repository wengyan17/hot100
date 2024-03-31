package main

import "fmt"

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

func main() {
	nums := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}
