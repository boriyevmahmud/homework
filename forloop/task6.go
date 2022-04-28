package main

import "fmt"

func main() {
	var a int
	for {
		fmt.Scan(&a)
		if a >= 10 && a <= 100 {
			fmt.Println(a)
		} else if a > 100 {
			break
		}
	}
}