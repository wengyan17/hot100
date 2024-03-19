package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	//fmt.Println("nums:", nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[l] + nums[r]
			//fmt.Println("target:", target, "sum:", sum, "l, r:", l, r)
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				res := []int{nums[i], nums[l], nums[r]}
				//fmt.Println("res", res)
				ans = append(ans, res)
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return ans
}

func main() {
	nums := []int{1, -1, -1, 0}
	fmt.Println(threeSum(nums))
}
