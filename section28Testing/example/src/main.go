package main

import (
	"acdc"
	"fmt"
)

func main() {
	fmt.Println(acdc.Sum([]int{1, 2, 3, 4, 5}...))

}
