package main

import "fmt"

func findAnagrams(s string, p string) (ans []int) {
	ss, sp := len(s), len(p)
	if sp > ss {
		return
	}
	count := [26]int{}

	for i := 0; i < sp; i++ {
		count[s[i]-'a']++
		count[p[i]-'a']--
	}

	diff := 0
	for _, i := range count {
		if i != 0 {
			diff++
		}
	}
	if diff == 0 {
		ans = append(ans, 0)
	}

	for i := sp; i < ss; i++ {
		//fmt.Println(count)
		count[s[i]-'a']++
		if count[s[i]-'a'] == 0 {
			diff--
		} else if count[s[i]-'a'] == 1 {
			diff++
		}
		//fmt.Println(diff)

		count[s[i-sp]-'a']--
		if count[s[i-sp]-'a'] == 0 {
			diff--
		} else if count[s[i-sp]-'a'] == -1 {
			diff++
		}
		//fmt.Println(diff)

		if diff == 0 {
			ans = append(ans, i-sp+1)
		}
	}
	return
}

func main() {
	s, p := "bpaa", "aa"
	fmt.Println(findAnagrams(s, p))
}
