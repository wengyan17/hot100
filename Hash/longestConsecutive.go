package main

import "fmt"

func longestConsecutive(nums []int) int {
	var ans = 0
	var mp = map[int]int{}
	for _, i := range nums {
		mp[i]++
	}
	for _, i := range nums {
		if mp[i-1] > 0 {
			continue
		}
		if mp[i] > 1 {
			mp[i]--
			continue
		}

		j := i + 1
		for {
			if mp[j] > 0 {
				j++
			} else {
				break
			}
		}
		ans = max(ans, j-i)
	}

	return ans
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
