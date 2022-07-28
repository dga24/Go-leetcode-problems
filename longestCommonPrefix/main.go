package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"dog","racecar","car"}
	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	for i := 0; i<len(strs); i++{
		if len(strings.Split(strs[0],""))==0 {
			fmt.Println("There is no common prefix among the input strings.")
					return ""
		}
		for{
			if strings.Compare(strs[0],strs[i])<=0{
				if strings.HasPrefix(strs[i],strs[0]){
					break
				}
			}else{
				if len(strings.Split(strs[0], ""))==0{
					break
				}
			}	
			strs[0] = strs[0][:len(strs[0])-1]		
		}
	}
	return strs[0]
}