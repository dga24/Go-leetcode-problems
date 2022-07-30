package main

import (
	"fmt"
	"strings"
)

// Given a sequence of moves for a robot. Check if the sequence is circular or not.

// A sequence of moves is circular if the first and last positions of the robot are the same. A move can be one of the following :
// G - Go one unit
// L - Turn left
// R - Turn right

// Example 1:
// Input: path = "GLGLGLG"
// Output: "Circular"
// Explanation: If we start form
// (0,0) in a plane then we will
// back to (0,0) by the end of the
// sequence.

// Example 2:
// Input: path = "GGGGL"
// Output: "Not Circular"
// Explanation: We can't return to
// same place at the end of the path.

// Your Task:
// You don't need to read input or print anything. Your task is to complete the function isCircular() which takes the string path as input and returns "Circular" if the path is circular else returns "Not Circular".




func main(){
	fmt.Println(isCircular("GLGLGLG"))
}

func isCircular(path string)  bool{
	slice := strings.Split(path,"")
	v := []int{0,0}
	for i:=0; i<len(slice);i++{
		if slice[i]=="G"{
		v[0]+=1
	  }else if slice[i]=="L"{
		v[0],v[1]=-v[1],v[0]
	  }else if slice[i]=="R"{
		v[0],v[1]=-v[1],v[0]
	  }
	}
	if v[1]==0 && v[0]==0{
	return true
	}
	return false
} 

