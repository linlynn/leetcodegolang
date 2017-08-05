// checkrecords.go
package main

import (
	"fmt"
)

func checkRecord(s string) bool {
	countA := 0
	continuousL := 0
	for _, j := range s {
		if j == 'A' {
			countA = countA + 1
			continuousL = 0
		} else if j == 'L' {
			continuousL = continuousL + 1
		} else {
			continuousL = 0
		}
		if countA > 1 || continuousL > 2 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println("Hello World!")
	x := checkRecord("LPLPAL")
	fmt.Println(x)
}
