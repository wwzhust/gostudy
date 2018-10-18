/*
	Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
*/

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "-12345894323.12345"
	var sign string
	var str []string

	if s[0] == '-' || s[0] == '+' {
		sign = string(s[0])
		str = strings.Split(s[1:], ".")
	} else {
		str = strings.Split(s[0:], ".")
	}

	fs := sign + commas(str[0]) + "." + str[1]
	fmt.Println(fs)
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
