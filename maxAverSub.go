// maxAverSub.go
package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	length := len(nums)
	var avg float64 = math.MinInt64
	var currSum float64 = 0
	var i int
	for i = 0; i < k; i++ {
		currSum = currSum + float64(nums[i])
	}
	avg = math.Max(avg, currSum/float64(k))
	for j := i; j < length; j++ {
		currSum = currSum - float64(nums[j-k]) + float64(nums[j])
		avg = math.Max(avg, currSum/float64(k))
	}
	return avg
}
func main() {
	x := []int{5}
	k := 1
	z := findMaxAverage(x, k)
	fmt.Println(z)
}
