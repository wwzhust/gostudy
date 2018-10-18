/*
Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of str ing
conc atenation.
*/

package main

import (
	"bytes"
	"fmt"
	//"strings"
)

func main() {
	s := "12345678"
	fmt.Println(commas(s))
}

func commas(s string) string {
	b3 := make([]byte, 3)
	buf := bytes.NewBufferString(reverseStr(s))
	var sr string
	for {
		if buf.Len() <= 3 {
			sr += buf.String()
			break
		}
		buf.Read(b3)
		sr += string(b3[:]) + ","
	}
	return reverseStr(sr)
}

func reverseStr(str string) string {
	b := []byte(str)
	l := len(str)
	li := l - 1
	s := make([]byte, l)
	for i := 0; i < l; i++ {
		s[i] = b[li-i]
	}

	return string(s)
}
