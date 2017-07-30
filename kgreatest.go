// kgreatest.go
package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}

func main() {
	fmt.Println("Hello World!")
	var test []int = []int{3, 2, 1, 5, 6, 4}
	var k int = 2
	y := findKthLargest(test, k)
	fmt.Println(y)
}
