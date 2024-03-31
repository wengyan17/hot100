package main

import "fmt"

func productExceptSelf(nums []int) (ans []int) {
	sum := 1
	numsZore := 0
	for _, i := range nums {
		if i == 0 {
			numsZore++
		} else {
			sum *= i
		}
	}

	if numsZore > 1 {
		for i := 0; i < len(nums); i++ {
			ans = append(ans, 0)
		}
	} else if numsZore == 1 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				ans = append(ans, sum)
			} else {
				ans = append(ans, 0)
			}
		}
	} else if numsZore == 0 {
		for _, i := range nums {
			ans = append(ans, sum/i)
		}
	}

	return
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
