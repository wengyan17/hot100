package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	n := len(intervals)
	for i := 1; i < n; i++ {
		last := len(ans) - 1
		if intervals[i][0] <= ans[last][1] {
			ans[last][1] = max(ans[last][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
