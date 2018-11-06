/*
	Exercise 4.4: Write a version of rotate that operates in a single pass.
*/
// Rev reverses a slice.
package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	ra := rotate(a[:], 2)
	fmt.Println("a: ", a)
	fmt.Println("ra: ", ra)

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	s = rotate(s, 2)
	fmt.Println("s: ", s) // "[2 3 4 5 0 1]"
	//!-slice

}

func rotate(s []int, pos int) []int {
	tmp := make([]int, len(s))
	copy(tmp[:len(s)-pos], s[pos:len(s)])
	copy(tmp[len(s)-pos:], s[:pos])
	fmt.Println("in rotate : ", s)
	return tmp
}
