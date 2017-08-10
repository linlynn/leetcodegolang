// distributeCandies.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func distributeCandies(candies []int) int {
	sort.Ints(candies)
	mapc := make(map[int]int)
	for _, j := range candies {
		mapc[j]++
	}
	lenm := len(mapc)
	return int(math.Min(float64(len(candies)/2), float64(lenm)))
}
func main() {
	fmt.Println("Hello World!")
	x := []int{1, 1, 2, 3}
	y := distributeCandies(x)
	fmt.Println((y))
}
