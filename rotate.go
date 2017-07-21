// rotate.py
package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	l := len(nums)
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		if (k + i) < l {
			ret[i+k] = nums[i]
		} else {
			ret[(i+k)%l] = nums[i]
		}

	}
	for i := 0; i < l; i++ {
		nums[i] = ret[i]
	}

}

func main() {
	var n []int = []int{2, 3}
	var k int = 1
	rotate(n, k)
}
