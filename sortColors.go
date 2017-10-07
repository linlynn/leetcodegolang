package main

import "fmt"


func sortColors(nums []int)  {
	nmap := make(map[int]int)
	for  _, j := range nums {
		nmap[j]++
	}
	j := 0
	for i := 0 ; i < 3 ; i++ {
		tempC := nmap[i]
		for x:= 0 ; x < tempC ; x ++ {
			nums[j]= i
			j++
		}
	}

}
func main() {
	nums := []int{0,1,2,1,2,1,0,0,0}
	sortColors(nums)
	fmt.Println(nums)
}
