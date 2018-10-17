/*
Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of str ing
conc atenation.
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "12345678"
	fmt.Println(commas(s))
}

func commas(s string) string {
	b3 := make([]byte, 3)
	buf := bytes.NewBufferString(s)
	var sr string
	for {
		if buf.Len() <= 3 {
			sr += buf.String()
			break
		}
		buf.Read(b3)
		sr += string(b3[:]) + ","
	}
	return sr
}
