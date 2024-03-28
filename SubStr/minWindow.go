package main

import (
	"fmt"
	"math"
)

func check(mps, mpt map[byte]int) bool {
	for k, v := range mpt {
		if mps[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) (ans string) {
	mps := map[byte]int{}
	mpt := map[byte]int{}

	for i := 0; i < len(t); i++ {
		mpt[t[i]]++
	}

	ansl, ansr := 0, math.MaxInt32

	for l, r := 0, 0; r < len(s); r++ {
		mps[s[r]]++

		for mpt[s[l]] == 0 {
			mps[s[l]]--
			if l < len(s)-1 {
				l++
			} else {
				break
			}
		}

		//for k, v := range mps {
		//	fmt.Println(k, " ", v)
		//}
		//fmt.Println(s[l:r+1], " ", check(mps, mpt))
		for check(mps, mpt) {
			if (r - l) < (ansr - ansl) {
				ansr, ansl = r, l
			}
			mps[s[l]]--
			l++
		}
	}

	if ansr == math.MaxInt32 {
		return ""
	}
	ans = s[ansl : ansr+1]
	return
}

func main() {
	s := "bba"
	t := "ab"

	fmt.Println(minWindow(s, t))
}
