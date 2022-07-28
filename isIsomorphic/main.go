package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s := "badc" 
	t := "baba"
	fmt.Println(isIsomorphic(s,t))

}

func isIsomorphic(s string, t string) bool {
	sls := strings.Split(s, "")
	slt := strings.Split(t, "")
	ok := false
	for i := 0; i < len(sls); i++ {	
		for j := 0; j < len(sls); j++ {
			if sls[j] == sls[i]{
				if slt[j] == slt[i]{
					ok = true
				}else{
					ok = false
				}
			}
			if slt[j] == slt[i]{
				if sls[j] == sls[i]{
					ok = true
				}else{
					ok = false
				}
			}
			if !ok{
				return false
			}
		}	
	}
	return true

}


func isIsomorphicOld(s string, t string) bool {
	sls := strings.Split(s, "")
	slt := strings.Split(t, "")
	if len(sls) != len(slt){
		return false
	}
	mapas := make([]int,len(sls))
	mapat := make([]int,len(sls))
	for i := 0; i < len(sls); i++ {
		for j := 0; j < len(sls); j++ {
			if sls[j] == sls[i]{
				mapas[j]=1
			}else{
				mapas[j]=0
			}
			if slt[j] == slt[i]{
				mapat[j]=1
			}else{
				mapat[j]=0
			}
		}
		if !reflect.DeepEqual(mapas,mapat){
			return false
		}
	}
	return true
}
	
