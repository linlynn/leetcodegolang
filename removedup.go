// removedup.go
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	newtail := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[newtail] {
			newtail = newtail + 1
			nums[newtail] = nums[i]
		}
	}
	return newtail + 1

}

func main() {
	fmt.Println("Hello World!")
	x := []int{1, 2, 2, 3}
	y := removeDuplicates(x)
	fmt.Println(y)
}
