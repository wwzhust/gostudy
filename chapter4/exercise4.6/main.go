// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	runes := []rune("Ooh, that  is sooo   cooool, 哈哈  呵呵  .")
	rs := delDupSpace(runes)
	fmt.Println("after delDup: ", string(rs))
}

func delDupSpace(r []rune) []rune {
	l := len(r)
	for i := 0; i < l-1; i++ {
		if unicode.IsSpace(r[i]) && r[i] == r[i+1] {
			copy(r[i:], r[i+1:])
			l--
			i--
		}
	}
	return r[:l]
}
