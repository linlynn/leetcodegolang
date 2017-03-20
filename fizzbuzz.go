// fizzbuzz.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	x := fizzBuzz(15)
	for _, j := range x {
		fmt.Printf("value is %s\n", j)
	}
}

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return res
}
