// hamming.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	x := hammingDistance(1, 4)
	fmt.Printf("distance is %d", x)
}

func hammingDistance(x int, y int) int {
	var result int = 0
	var xslic, yslic [32]int
	for i, index := x, 0; ; index = index + 1 {
		xslic[index] = i % 2
		i = i / 2
		if i == 0 {
			index = index + 1
			xslic[index] = i % 2
			break
		}
	}
	for j, index := y, 0; ; index = index + 1 {
		yslic[index] = j % 2
		j = j / 2
		if j == 0 {
			index = index + 1
			yslic[index] = j % 2
			break
		}
	}
	for a, b := range yslic {
		fmt.Printf("a and b is %d,%d\n", a, b)
	}
	for i, j := range xslic {
		if j != yslic[i] {
			result++
		}
	}
	return result
}
