// validpalindrome.go
/*

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

For example,
"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.


*/
package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	var l int = 0
	var r int = len(s) - 1
	al := "abcdefghijklmnopqrstuvwxyz0123456789"
	for l < r {
		for l < r {
			if strings.Contains(al, strings.ToLower(string(s[l]))) == false {
				l = l + 1
				fmt.Println("helloaaaa")
			} else {
				break
			}
		}
		for l < r {
			if strings.Contains(al, strings.ToLower(string(s[r]))) == false {
				r = r - 1
			} else {
				break
			}
		}
		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		} else {
			l = l + 1
			r = r - 1
		}
	}
	return true
}

func main() {
	x := "0P"
	y := isPalindrome(x)
	fmt.Println(y)
}
