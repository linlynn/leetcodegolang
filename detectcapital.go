// detectcapital.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	x := detectCapitalUse("GG")
	fmt.Println(x)
}

func detectCapitalUse(word string) bool {
	cap := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lenw := len(word)
	num := 0
	for i := 0; i < lenw; i++ {
		if strings.ContainsAny(cap, string(word[i])) {
			num++
		}
	}
	if num == lenw {
		return true
	} else if num == 1 && strings.ContainsAny(cap, string(word[0])) {
		return true
	} else {
		return false
	}

}
