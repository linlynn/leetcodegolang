// judgeroute.go
package main

import (
	"fmt"
	"strings"
)

func judgeCircle(moves string) bool {
	return strings.Count(moves, "L") == strings.Count(moves, "R") && strings.Count(moves, "U") == strings.Count(moves, "D")
}

func main() {
	fmt.Println("Hello World!")
	var x string = "LRD"
	fmt.Println(judgeCircle(x))
}
