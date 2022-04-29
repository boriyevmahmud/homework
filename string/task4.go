package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	for i := 1; i < len(a); i += 2 {
		fmt.Print(string(a[i]))
	}

}
