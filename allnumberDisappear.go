package main

import "fmt"

func main() {
	fmt.Printf("hello,world\n")
	x := []int{4, 3, 2, 7, 8, 2, 3, 1}
	y := findDisappearedNumbers(x)
	for _, value := range y {
		fmt.Printf("value is %d\n", value)
	}
}

func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	res := make([]int, 0)
	nhash := make(map[int]int)
	for _, j := range nums {
		nhash[j]++
	}
	for i := 1; i <= length; i++ {
		if _, ok := nhash[i]; !ok {
			res = append(res, i)
			fmt.Printf("====i is %d\n", i)
		}
	}
	return res

}
