/*
	Exercise 3.13: Write const declarations for KB, MB, up through YB as compactly as you can
*/

package main

import (
	"fmt"
	//"math"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
)

func main() {
	fmt.Println(KB, MB, YB)
}
