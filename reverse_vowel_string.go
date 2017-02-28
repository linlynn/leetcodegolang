package main

import "fmt"
import "strings"

func reverseVowels(s string) string {
	vo := "aeiouAEIOU"
	lens := len(s)
	r := []rune(s)
	for i, j := 0, lens-1; i < j; i, j = i+1, j-1 {
		for temp := i; temp < j; temp++ {
			if !strings.ContainsRune(vo, rune(s[temp])) {
				i++
			} else {
				break
			}
		}
		for tempj := j; tempj > i; tempj-- {
			if !strings.ContainsRune(vo, rune(s[tempj])) {
				j--
			} else {
				break
			}
		}
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main() {
	var test string = "leetcode"
	res := reverseVowels(test)
	fmt.Printf("reversed string is %s\n", res)
}
