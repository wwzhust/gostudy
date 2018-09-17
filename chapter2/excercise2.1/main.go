package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	type TestConv struct {
		ctemp tempconv.Celsius
		ftemp tempconv.Fahrenheit
		ktemp tempconv.Kelvin
	}

	testconv := TestConv{32, 212, -40}

	fmt.Printf("testconv ctemp to F:%s \t ftemp to C:%s \t ktemp to C:%s \n", tempconv.CToF(testconv.ctemp).String(), tempconv.FToC(testconv.ftemp).String(), tempconv.KToC(testconv.ktemp).String())
}
