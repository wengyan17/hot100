package main

import (
	"fmt"
	"sort"
)

//func twoSum(nums []int, target int) []int {
//	mp := map[int]int{}
//	for i := 0; i < len(nums); i++ {
//		mp[nums[i]] = i
//	}
//
//	sort.Ints(nums)
//	var arr []int
//	for i := 0; i < len(nums); i++ {
//		ans := check(target-nums[i], nums)
//		if nums[i]+nums[ans] == target {
//			arr = append(arr, mp[nums[i]], mp[nums[ans]])
//			break
//		}
//	}
//
//	sort.Ints(arr)
//	return arr
//}
//
//func check(x int, nums []int) int {
//	l := 0
//	r := len(nums)
//	for {
//		if l+1 == r {
//			break
//		}
//		mid := (l + r) / 2
//		if nums[mid] <= x {
//			l = mid
//		} else {
//			r = mid
//		}
//	}
//	return l
//}

func twoSum(nums []int, target int) []int {
	mp := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = append(mp[nums[i]], i)
	}

	var ans []int
	flag := false
	for i := 0; i < len(nums); i++ {
		if flag {
			break
		}
		fst := nums[i]
		for _, num := range mp[target-fst] {
			if num != i {
				flag = true
				ans = append(ans, i, num)
				break
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
