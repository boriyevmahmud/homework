package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	m := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m[i])
	}
	fmt.Println(m[3])
}
