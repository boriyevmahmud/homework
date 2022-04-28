package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for _, elem := range arr {
		if elem > 0 {
			sum += 1
		}
	}
	fmt.Print(sum)
}
