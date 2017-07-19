// findalldup.go
package main

import (
	"fmt"
)

// input [4 3 2 7 8 2 3 1]
func findDuplicates(nums []int) []int {
	var ret []int
	hashn := make(map[int]int)
	for _, j := range nums {
		hashn[j]++
	}
	for i, j := range hashn {
		if j == 2 {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {
	fmt.Println("Hello World!")
	x := []int{4, 3, 2, 7, 8, 2, 3, 1}
	y := findDuplicates(x)
	for _, j := range y {
		fmt.Println(j)
	}
}
