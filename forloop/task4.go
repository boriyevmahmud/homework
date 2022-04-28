package main

import "fmt"

func main() {
	var a, max, count int
	for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
		if a > max{
			max = a
			count=0
		} 
		if a==max {
			count+=1
		}
	}
	fmt.Println(count)
}