package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	//"time"
)

//exercise 1.4
// Modify dup2 to print the names of all files in which each duplicated line occurs
func main() {
	counts := make(map[string]int)
	fcounts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fcounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fcounts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fcounts[line])
		}
	}
}

func inarray(fname string, names []string) bool {
	for _, s := range names {
		if s == fname {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, fcounts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !inarray(f.Name(), fcounts[input.Text()]) {
			fcounts[input.Text()] = append(fcounts[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
