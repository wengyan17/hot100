package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, i := range strs {
		sli := strings.Split(i, "")
		sort.Strings(sli)
		sortcli := strings.Join(sli, "")
		mp[sortcli] = append(mp[sortcli], i)
	}

	var ans [][]string
	for _, i := range mp {
		ans = append(ans, i)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)
}
