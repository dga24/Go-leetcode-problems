package main

import (
	"fmt"
	"sort"
)

func main(){
	arr := []int32{3,7,5,6,2}
	fmt.Println(minimalHeaviestSetA(arr))
}

func minimalHeaviestSetA(arr []int32) []int32 {
	sort.Slice(arr, func(i int, j int) bool {
		return arr[i] > arr[j]
	})
	sumA := 0
	total := 0
	size := 0
	var res []int32
	for i := 0; i < len(arr); i++ {
		total += int(arr[i])
	}
	for i := 0; i < len(arr); i++ {
		sumA += int(arr[i])
		if sumA > (total - sumA) {
			size = i
			break
		}
	}
	for i := 0; i <= size; i++ {
		res = append(res, arr[i])
	}
	sort.Slice(res, func(i int, j int) bool {
		return res[i] < res[j]
	})
	return res

}