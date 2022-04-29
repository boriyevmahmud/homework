package main

import (
	"fmt"
)

func main() {
	var x, s string
	var count string
	fmt.Scan(&s, &s)
	for i := 0; i < len(s); i++ {
		if x[i] == s[i] {
			count += 1
		}
	}
	fmt.Print(count)
}
