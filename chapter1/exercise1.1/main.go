package main

import (
	"fmt"
	"os"
)

//exercise 1.1
func main() {
	var s, sep string
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}