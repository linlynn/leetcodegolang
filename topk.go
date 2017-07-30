// topk.go
package main

import (
	"fmt"
	"sort"
)

//input [1,2] and k = 2

// hashs {1:1,2:1}

//

func topKFrequent(nums []int, k int) []int {
	hashn := make(map[int]int)
	for _, j := range nums {
		hashn[j]++
	}
	n := map[int][]int{}
	var a []int
	var res []int
	for i, j := range hashn {
		n[j] = append(n[j], i)
	}

	for i, _ := range n {
		a = append(a, i)
	}
	// a [1,2]
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// a [2,1]
	temp := 0

	for _, j := range a {
		if temp+len(n[j]) <= k {
			for _, t := range n[j] {
				res = append(res, t)
				temp = temp + 1
				fmt.Println("temp is")
				fmt.Println(temp)
				fmt.Println(res)
			}
		} else {
			for _, tt := range n[j] {
				if temp < k {
					res = append(res, tt)
					temp = temp + 1
				} else {
					break
					return res
				}
			}
		}
	}
	return res

}

func main() {
	var num []int = []int{1, 1, 1, 2, 2, 3}
	k := 2
	x := topKFrequent(num, k)
	fmt.Println(x)
}
