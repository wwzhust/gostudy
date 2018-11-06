/*
	Exercise 4.5 : Write an in-place function to eliminate adjacent duplicates in a []string slice
*/
package main

import (
	"fmt"
)

func main() {
	s := "Ooh, that is sooo cooool, hahaa"
	rs := delDup([]byte(s))
	fmt.Println("after delDup: ", rs)

	str := []string{"what", "which", "which", "which", "why", "why", "how", "when", "who"}
	rstr := delDupStr1(str)
	fmt.Println("after delDupStr: ", rstr)
}

func delDup(s []byte) string {
	l := len(s)
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:l])
			l--
			i--
		}
	}
	return string(s[:l])
}

func delDupStr(s []string) []string {
	l := len(s)
	for i := 0; i < l-1; i++ {
		//fmt.Println(s[i], s[i+1])
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			l--
			i--
		}
	}
	return s[:l]
}

func delDupStr1(s []string) []string {
	l := len(s)
	for i := 0; i < l-1; i++ {
		//fmt.Println(s[i], s[i+1])
		if s[i] == s[i+1] {
			for j := i; j < l-1; j++ {
				s[j] = s[j+1]
			}
			l--
			i--
		}
	}
	return s[:l]
}
