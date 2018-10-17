/*
	Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,that is, they contain the same letters in a different order.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please input two references!!!")
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	if isAnagram(str1, str2) {
		fmt.Printf("%s and %s are anagrams\n", str1, str2)
	} else {
		fmt.Printf("%s and %s are NOT anagrams\n", str1, str2)
	}
}

func isAnagram(str1, str2 string) bool {
	strb1 := bytsToString(str1)
	strb2 := bytsToString(str2)

	sort.Strings(strb1)
	sort.Strings(strb2)

	fstr1 := strings.Join(strb1, "")
	fstr2 := strings.Join(strb2, "")

	fmt.Println(fstr1)
	fmt.Println(fstr2)
	if fstr1 == fstr2 {
		return true
	}

	return false
}

func bytsToString(str string) []string {
	/*var b = []byte(str)
	rstr := make([]string, len(b))
	for i, _ := range b {
		rstr[i] = string(b[i])
	}
	return rstr*/
	return strings.Split(str, "")
}
