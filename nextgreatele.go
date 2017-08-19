// nextgreatele.go
package main

import (
	"fmt"
)

/*
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
    For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
*/

func nextGreaterElement(findNums []int, nums []int) []int {
	res := []int{}
	var k int = 0
	for i := 0; i < len(findNums); i++ {
		temp := findNums[i]
		for j := 0; j < len(nums); j++ {
			if nums[j] == temp {
				for k = j + 1; k < len(nums); k++ {
					if nums[k] > temp {
						res = append(res, nums[k])
						break
					}
				}
				if k >= len(nums) {
					res = append(res, -1)
				}
				break
			} else {
				continue
			}
		}
	}
	return res
}

func main() {
	fmt.Println("Hello World!")
	x := []int{1, 3, 5, 2, 4}
	y := []int{6, 5, 4, 3, 2, 1, 7}
	z := nextGreaterElement(x, y)
	fmt.Println(z)
}
