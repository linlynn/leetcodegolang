// keyboard.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	y := []string{"hello", "hlaska", "dad", "peace"}
	x := findWords(y)
	for _, j := range x {
		fmt.Printf("string is %s\n", j)
	}
}

//Input: ["Hello", "Alaska", "Dad", "Peace"]
//Output: ["Alaska", "Dad"]
//
func findWords(words []string) []string {
	var row1 string = "qwertyuiop"
	var row2 string = "asdfghjkl"
	var row3 string = "zxcvbnm"
	var result []string
	var startrow1, startrow2, startrow3 int
	for _, res := range words {
		reslow := strings.ToLower(res)
		for _, y := range reslow {
			if strings.Contains(row1, string(y)) == true {
				startrow1++
			} else if strings.Contains(row2, string(y)) == true {
				startrow2++
			} else if strings.Contains(row3, string(y)) == true {
				startrow3++
			} else {
				fmt.Println("nothing")
			}
		}
		if len(res) == startrow1 || len(res) == startrow2 || len(res) == startrow3 {
			result = append(result, res)
		}
		startrow1, startrow2, startrow3 = 0, 0, 0
	}
	return result
}
