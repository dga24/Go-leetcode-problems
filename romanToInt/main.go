package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	symbolMap := make(map[string]int)
	symbolMap["M"] = 1000
	symbolMap["D"] = 500
	symbolMap["C"] = 100
	symbolMap["L"] = 50
	symbolMap["X"] = 10
	symbolMap["V"] = 5
	symbolMap["I"] = 1

	res := 0
	input := strings.Split(s, "")

	for i :=0; i < len(input);i++{
		if i>0 && (symbolMap[input[i-1]] < symbolMap[input[i]]) {
			res -= symbolMap[input[i-1]]
			res += (symbolMap[input[i]]-symbolMap[input[i-1]]) 
		}else{
			res += symbolMap[input[i]]
		}
	}
	return res
}