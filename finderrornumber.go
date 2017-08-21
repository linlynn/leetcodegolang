// finderrornumber.go
package main

import (
	"fmt"
)

/*

Example 1:
Input: nums = [1,2,2,4]
Output: [2,3]
*/

func findErrorNums(nums []int) []int {
	res := []int{}
	hash1 := make(map[int]int)
	for _, j := range nums {
		_, ok := hash1[j]
		if ok {
			res = append(res, j)
		} else {
			hash1[j] = j
		}
	}
	for i := 1; i <= len(nums); i++ {
		_, ok := hash1[i]
		if !ok {
			res = append(res, i)
			return res
		}
	}
	return res

}

func main() {
	fmt.Println("Hello World!")
	x := []int{1, 2, 2, 4}
	y := findErrorNums(x)
	fmt.Println(y)
}
