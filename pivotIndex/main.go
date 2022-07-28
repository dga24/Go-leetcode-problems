package main

import "fmt"

func main() {
	nums := []int{2,1,-1}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j<i {
				sum += nums[j]
			}else if j>i{
				sum -= nums[j]
			}
		}
		if sum==0{
			return i
		}
		sum= 0
	}
	return -1
}

