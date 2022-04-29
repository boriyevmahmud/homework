package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var b []rune = []rune(text)
	if unicode.IsUpper(b[0]) && string(b[len(b)-1]) == "." {
		fmt.Println("Right")
	} else {giz
		fmt.Println("Wrong")
	}
}
