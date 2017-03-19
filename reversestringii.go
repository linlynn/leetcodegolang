// reversestringii
package main

import (
	"fmt"
)

func main() {
	x := reverseStr("abcdeg", 2)
	fmt.Println(x)
}

func reverseStr(s string, k int) string {

	lens := len(s)
	r := []rune(s)
	if lens < k {
		for i, j := 0, lens-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}
	if lens >= k && lens < 2*k {
		for i, j := 0, k-1; i <= j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}
	for i := 0; i < lens; i = i + 2*k {
		if i+2*k-1 < lens {
			for x, y := i, i+k-1; x < y; x, y = x+1, y-1 {
				r[x], r[y] = r[y], r[x]
			}
			continue
		} else if lens-i < 2*k && lens-i >= k {
			for x, y := i, i+k-1; x < y; x, y = x+1, y-1 {
				r[x], r[y] = r[y], r[x]
			}
			continue
		} else {
			for x, y := i, lens-1; x < y; x, y = x+1, y-1 {
				r[x], r[y] = r[y], r[x]
			}
			continue
		}
	}
	return string(r)

}
