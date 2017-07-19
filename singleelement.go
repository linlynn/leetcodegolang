// singleelement.go
package main

import (
	"fmt"
)

func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		j := i + 1
		if j < len(nums) {
			if nums[i] != nums[j] {
				return nums[i]
			} else {
				i = i + 2
				continue
			}
		} else {
			return nums[i]
		}
	}
	return nums[0]
}

func main() {
	fmt.Println("Hello World!")
}
