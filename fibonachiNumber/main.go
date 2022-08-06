package main

import "fmt"

func main() {
	fmt.Println(fib(3))
}

func fib(n int) int {
	a ,b := 0, 1
	for i:= 0; i<n;i++{
		a,b = b, a+b
	}
	return a
}