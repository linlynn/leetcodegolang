// reversestringiii.go
package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseWords(s string) string {
	var r []rune
	words := strings.Split(s, " ")
	var temp []string
	for i := 0; i < len(words); i++ {
		temp = append(temp, words[i])
	}

	for i, j := range temp {
		x := reverseString(j)
		temp[i] = x

	}

	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp[i]); j++ {
			r = append(r, rune(temp[i][j]))
		}
		if i != len(temp)-1 {
			r = append(r, ' ')
		}
	}

	return string(r)
}

func main() {
	var s string = "hello world"
	res := reverseWords(s)
	fmt.Println(res)
}
