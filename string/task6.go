package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a string
	var b bool
	var temp rune
	fmt.Scan(&a)
	if len(a) >= 5 {
		for i := 0; i < len(a); i++ {
			temp = rune(a[i])
			if unicode.Is(unicode.Latin, temp) || unicode.IsDigit(temp) {
				b = true	

			} else {
				b = false
				break
			}
		}
		if b == true {
			fmt.Println("Ok")
		} else {
			fmt.Println("Wrong password")
		}

	} else {
		fmt.Println("Wrong password")
	}
}
