// addstring.go
package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	var res []rune
	var a, b, carry, sum int
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		if i >= 0 {
			a = int(num1[i] - '0')
			i--
		} else {
			a = 0
		}
		if j >= 0 {
			b = int(num2[j] - '0')
			j--
		} else {
			b = 0
		}
		sum = a + b + carry
		carry = sum / 10
		x := rune(sum%10 + '0')
		res = append([]rune{x}, res...)

	}
	if carry > 0 {
		res = append([]rune{'1'}, res...)
	}
	return string(res)

}
func main() {
	x := addStrings("1", "19")
	fmt.Printf("add string is %s\n", x)
}
