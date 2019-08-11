package main

import (
	"fmt"
	"pi"
	"runtime"
)

func main() {

	r := pi.CalculatePiByArray(1000)
	fmt.Println(r)

	r = pi.CalculatePiByFanInOutChannel(100000, runtime.NumCPU())
	fmt.Println(r)
}
