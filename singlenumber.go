// singlenumber
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	num := []int{1, 1, 8, 3, 2, 3, 2}

	x := singleNumber(num)
	fmt.Printf("x is %d\n", x)
}

func singleNumber(nums []int) int {
	mymap := make(map[int]int)
	for _, j := range nums {
		mymap[j]++
	}
	for i, j := range mymap {
		if j == 1 {
			return i
		}
	}
	return 0
}
