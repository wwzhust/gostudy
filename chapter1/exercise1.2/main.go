package main

import (
	"fmt"
	"os"
)


//exercise1.2

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, " : ", arg)
	}
}
