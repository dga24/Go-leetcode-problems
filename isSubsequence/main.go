package main

import (
	"fmt"
)

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s,t))

}

func isSubsequence(s string, t string) bool {
	if len(s)==0{
		return true
	}
	iS, iT := 0, 0
	for iT < len(t) {
		if t[iT] == s[iS] {
			iS++
			if iS == len(s) {
				return true
			}
		}
		iT++
	}

	return false
}