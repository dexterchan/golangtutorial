package main

import "fmt"

func main() {
	//i := []int{2, 3}

	fmt.Println("2+3=", mySum(2, 3))

	fmt.Println("2+3=", mySum([]int{2, 3}...))
}
func mySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
