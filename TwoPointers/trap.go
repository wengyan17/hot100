package main

import (
	"fmt"
)

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

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
