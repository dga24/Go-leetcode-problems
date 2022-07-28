package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println((runningSum(nums)))
}

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	for i, _ := range nums {
		for j := 0; j <= i; j++ {
			res[i] += nums[j]
		}
	}
	return res
}