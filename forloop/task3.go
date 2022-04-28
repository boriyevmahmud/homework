package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Scan(&b)
		if b > 9 && b < 100 {
			if b%8 == 0 {
				sum += b
			}
		}
	}
	fmt.Println(sum)
}