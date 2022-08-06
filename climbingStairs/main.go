package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
    path := []int{1,1}
    for i:=2;i<=n;i++ {
        path = append(path, path[i-1] + path[i-2])
    }
    return path[n]
}






