package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//exercise 1.1
/*
func main() {
	var s, sep string
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
*/

//exercise1.2
/*
func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, " : ", arg)
	}
}
*/

//exercise 1.3
func main() {
	var s, sep string
	start := time.Now()
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	ect_time := end.Sub(start)
	fmt.Println("+: ", ect_time) // for arg:0...5000 ; shown "+:  14.0743ms"

	start = time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	end = time.Now()
	ect_time = end.Sub(start)
	fmt.Print("join: ", ect_time) // for arg:0...5000 ; shown "join: 0s"
}
