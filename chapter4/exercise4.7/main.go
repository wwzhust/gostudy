/*
	Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a
	UTF-8-encoded string , in place. Can you do it without allocating new memory?
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	runes := []byte("一二三四五六七八九十")
	rs := reverseBytes(runes)
	fmt.Println(string(rs))
}

func reverseBytes(a []byte) []byte {
	if utf8.RuneCount(a) == 1 {
		return a
	}
	_, s := utf8.DecodeRune(a)
	return append(reverseBytes(a[s:]), a[:s]...)
}
