// strstr.go
package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	lenh := len(haystack)
	lenn := len(needle)
	for i := 0; i < lenh-lenn+1; i++ {
		if haystack[i:i+lenn] == needle {
			return i
		}
	}
	return -1

}

func main() {
	fmt.Println("Hello World!")
	x := "abcde"
	y := "cde"
	z := strStr(x, y)
	fmt.Println(z)
}
