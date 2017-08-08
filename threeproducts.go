// threeproducts.go
package main

import (
	"fmt"
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return int(math.Max(float64(nums[0]*nums[1]*nums[len(nums)-1]), float64(nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1])))
}
func main() {
	fmt.Println("Hello World!")
	var x []int = []int{-1, 2, 3, 4}
	y := maximumProduct(x)
	fmt.Println(y)
}
