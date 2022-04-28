package main

import "fmt"

func main() {
	array := [100]int{}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}

}
