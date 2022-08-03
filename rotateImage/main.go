package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix[0][2])
	fmt.Println(rotate(matrix))
}

func rotate(matrix [][]int) [][]int{
	n := len(matrix)
    for i := 0; i < (n+1)/2; i++ {
        for j := i; j < n-i-1; j++ {
            matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] = matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]            
        }
    }
	return matrix
}