// findUnsortedSubarray.go
package main

import (
	"fmt"
	"sort"
)

//Input: [2, 6, 4, 8, 10, 9, 15]
//Output: 5
//Explanation: You need to sort [6, 4, 8, 10, 9] in as

func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	var start, end int = 0, 0
	nums2 := make([]int, length)
	copy(nums2, nums)
	sort.Ints(nums2)
	for start < length && nums[start] == nums2[start] {
		start++
	}
	end = length - 1
	for end > start && nums[end] == nums2[end] {
		end--
	}
	return end - start + 1

}

func main() {
	x := []int{2, 6, 4, 8, 10, 9, 15}
	y := findUnsortedSubarray(x)
	fmt.Println(y)
}
