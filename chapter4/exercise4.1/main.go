/*
	Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
	(See PopCount from Sec tion 2.6.2.)
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func main() {
	//fmt.Println(PopCount(uint64(5)))
	c1 := sha256.Sum256([]byte("Helloworld1"))
	c2 := sha256.Sum256([]byte("Helloworld2"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	count := countDiffBits(c1, c2)
	fmt.Println(count)
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Println(i, " : ", pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
/*func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}*/

func countDiffBits(c1, c2 [32]uint8) int {
	var count uint8 = 0
	for i := uint8(0); i < uint8(32); i++ {
		x := c1[i] ^ c2[i] //get xor of c1 and c2,if a bit is different,then will be set
		count += pc[x]
	}
	return int(count)
}
