// test.go
package main

import (
	"fmt"
)

// s="ab"
// t="bae"

func findTheDifference(s string, t string) byte {
	hashs := make(map[byte]int)
	for _, j := range s {
		hashs[byte(j)]++
	}
	hasht := make(map[byte]int)
	for _, j := range t {
		hasht[byte(j)]++
	}

	for i, j := range hasht {
		_, ok := hashs[i]
		if !ok {
			return i
		} else {
			if j != hashs[i] {
				return i
			}
		}

	}
	fmt.Println("hello")
	return byte(0)
}

func main() {
	var s string = "a"
	var t string = "aa"
	x := findTheDifference(s, t)
	fmt.Printf("%c", x)
}
