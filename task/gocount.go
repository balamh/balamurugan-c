package main

import (
	"fmt"
)

func find(a []byte) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 'g' && a[i+1] == 'o' && (i+2==len(a) ||a[i+2] == ' '){
			count++
		}
	}
	return count
}
func main() {
	value := "hello go gog go hello go go"
	x := find([]byte(value))
	fmt.Println(x)
}
