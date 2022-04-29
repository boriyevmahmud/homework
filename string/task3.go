package main
import (
	"fmt"
	"strings"
)
func main() {
	var x, s string
	fmt.Scan(&x, &s)
	fmt.Print(strings.Index(x, s))
}
