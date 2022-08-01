package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "badc" 
	t := "baba"
	fmt.Println(isIsomorphic(s,t))

}
	

//this work on 41/43 tests
func isIsomorphic(s string, t string) bool {
	sls := strings.Split(s, "")
	slt := strings.Split(t, "")
	ok := false
	for i := 0; i < len(sls); i++ {	
		for j := 0; j < len(sls); j++ {
			if sls[j] == sls[i]{
				if slt[j] == slt[i]{
					ok = true
				}else{
					ok = false
				}
			}
			if slt[j] == slt[i]{
				if sls[j] == sls[i]{
					ok = true
				}else{
					ok = false
				}
			}
			if !ok{
				return false
			}
		}	
	}
	return true

}