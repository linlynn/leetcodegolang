// removelements.go
package main

import (
	"fmt"
)

// Example:
//Given input array nums = [3,2,2,3], val = 3

//Your function should return length = 2, with the first two elements of nums being 2.

func removeElement(nums []int, val int) int {
	newtail := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[newtail] = nums[i]
			newtail++
		}
	}
	return newtail

}

func main() {
	fmt.Println("Hello World!")
	x := []int{3, 2, 2, 3}
	y := removeElement(x, 3)
	fmt.Println(y)
}
