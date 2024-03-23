package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	ans := 0
	var mp = map[uint8]bool{}
	var str = ""
	for i := 0; i < len(s); i++ {
		if mp[s[i]] {
			for mp[s[i]] {
				mp[str[0]] = false
				str = str[1:]
			}
			str = str + string(s[i])
			mp[s[i]] = true
			ans = max(ans, len(str))
		} else {
			str = str + string(s[i])
			mp[s[i]] = true
			ans = max(ans, len(str))
		}
	}
	return ans
}

func main() {
	s := "aab"
	fmt.Println(lengthOfLongestSubstring(s))
}
