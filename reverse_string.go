// reverse_string
package main

import (
	"fmt"
)

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main() {
	res := reverseString("")
	fmt.Printf("reversed string is %s\n", res)
}
