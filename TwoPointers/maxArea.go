package main

import "fmt"

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

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
