package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums,9))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i!= j {
				if nums[i]+nums[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}