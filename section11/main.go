package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(x)
	for i := 0; i < len(x); i += 3 {
		fmt.Println(x[i : i+3])
	}

	xx := []string{"apple", "orange", "berry"}
	yy := []string{"red", "orange", "red"}
	m := [][]string{xx, yy}
	for _, nv := range m {
		//fmt.Println(ni)
		//ni = ni
		for ni1, nv2 := range nv {
			fmt.Println(ni1, nv2)
		}

	}
}
