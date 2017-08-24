// canPlaceFlower.go
package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	count := 0
	var next, prev int
	for i := 0; i < length && count < n; i++ {
		if flowerbed[i] == 0 {
			if i == length-1 {
				next = 0
			} else {
				next = flowerbed[i+1]
			}
			if i == 0 {
				prev = 0
			} else {
				prev = flowerbed[i-1]
			}
			if next == 0 && prev == 0 {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count == n
}

func main() {
	x := []int{1, 2, 0, 0, 0, 1}
	y := canPlaceFlowers(x, 2)
	fmt.Println(y)
}
