package main

import (
	"fmt"
	"sort"
)

func main() {
	parcels := []int32{int32(4), int32(2), int32(3), int32(4)}
	fmt.Println(getMinimumDays(parcels))
}

func getMinimumDays(parcels []int32) int32 {

	// Write your code here
	//the input is not correct
	fmt.Println("len:", len(parcels), "input:", parcels)

	sort.Slice(parcels, func(i, j int) bool { return parcels[i] < parcels[j] })
	pcls := make([]int, len(parcels))
	for i := 0; i < len(parcels); i++ {
		pcls[i] = int(parcels[i])
	}
	fmt.Println(pcls)
	min := pcls[0]
	var days int32
	days = int32(0)
	for {
		min = pcls[len(parcels)-1]
		for i := 0; i < len(pcls); i++ {
			if pcls[i] < min && pcls[i] > 0 {
				min = pcls[i]
				fmt.Println("min:", min)
			}
		}
		for i := 0; i < len(parcels); i++ {
			if pcls[i]>0{
                pcls[i]= pcls[i]-min
            }
		}
		days++
		if pcls[len(pcls)-1] == 0 {
			break
		}
	}

	// return int32(days)
	return days
}