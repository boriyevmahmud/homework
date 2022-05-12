package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// package main уже объявлен.
func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		sum += i
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
}
