package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	noZ := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[noZ], nums[r] = nums[r], nums[noZ]
			//fmt.Println(nums)
			noZ++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
