/*
	Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
words instead of lines.
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := "The temprature in Hawaii ranges from 55 degrees to 85 degrees,but the temprature in 深圳 在 5 到 38 度 之间"
	counts := make(map[string]int)

	rd := bufio.NewScanner(strings.NewReader(str))
	rd.Split(bufio.ScanWords)
	for rd.Scan() {
		counts[rd.Text()]++
	}

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
