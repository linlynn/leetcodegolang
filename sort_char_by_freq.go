// sort_char_by_freq.go
package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, j := range s {
		m[j]++
	}
	n := map[int][]rune{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	var ret []rune
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			for i := 0; i < k; i++ {
				ret = append(ret, s)
			}
		}
	}
	return string(ret)

}

func main() {
	s := "adccc"
	x := frequencySort(s)
	fmt.Println(x)
}
