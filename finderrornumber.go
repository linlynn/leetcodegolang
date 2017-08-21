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
		hash1[j]++
	}
	for i, j := range hash1 {
		if j == 2 {
			res = append(res, i)
		}
	}
	for i := 0; i < len(nums); i++ {
		_, ok := hash1[i+1]
		if !ok {
			res = append(res, i+1)
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
