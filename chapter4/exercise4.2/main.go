/*
	Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
	(See PopCount from Sec tion 2.6.2.)
*/
package main

import (
	"crypto/sha256"
	//"crypto/sha384"
	"crypto/sha512"
	"fmt"
	"os"
)

var pc [256]byte

func main() {
	if len(os.Args) == 0 || os.Args[1] == "256" {
		c1 := sha256.Sum256([]byte("X"))
		c2 := sha256.Sum256([]byte("x"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	} else if os.Args[1] == "512" {
		c1 := sha512.Sum512([]byte("X"))
		c2 := sha512.Sum512([]byte("x"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	} else if os.Args[1] == "384" {
		c1 := sha512.Sum384([]byte("X"))
		c2 := sha512.Sum384([]byte("x"))
		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	}
}
