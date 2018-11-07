/*
	Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
	using functions like unicode.IsLetter
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "The temprature in Hawaii ranges from 55 degrees to 85 degrees, 但是深圳在5到38度之间"
	counts := make(map[rune]int) // counts of Unicode characters
	lettercounts := make(map[rune]int)
	digitalcounts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	//in := bufio.NewReader(os.Stdin)
	in := bufio.NewReader(strings.NewReader(str))
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			lettercounts[r]++
		}
		if unicode.IsDigit(r) {
			digitalcounts[r]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("counts: %d\t%d\n", i, n)
		}
	}
	for c, n := range lettercounts {
		fmt.Printf("lettercounts : %q\t%d\n", c, n)
	}
	for c, n := range digitalcounts {
		fmt.Printf("digitalcounts: %q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
