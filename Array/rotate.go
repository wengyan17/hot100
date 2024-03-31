package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	newNums := append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	copy(nums, newNums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
