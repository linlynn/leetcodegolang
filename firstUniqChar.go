// firstUniqChar.go
package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	hash := make(map[rune]int, 26)
	lens := len(s)
	for i, _ := range s {
		hash[rune(s[i])]++
	}
	for i := 0; i < lens; i++ {
		if hash[rune(s[i])] == 1 {
			return i
		}
	}
	return -1
}
func main() {
	teststring := "loveleecode"
	index := firstUniqChar(teststring)
	fmt.Printf("string is %s and index is %d\n", teststring, index)
}
