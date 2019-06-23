package main

import (
	"fmt"
)

func main() {
	//simpleLoop()
	//whileLoop()
	//asciiloop()
	elseif()
	switchfunc()
}

func simpleLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func whileLoop() {
	x := 1
	for x < 10 {
		x++
	}
	fmt.Println(x)
}

func asciiloop() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#[2]x\t%#[2]U\n", i, i)
	}
}

func elseif() {
	x := 42
	y := 3
	if x == 40 {
		fmt.Println("our value is 40")
	} else if x == 42 || y == 2 {
		fmt.Println("our value was 42")
	} else {
		fmt.Println("other values")
	}

}

func switchfunc() {
	switch {
	case false:
		fmt.Println("this is false")
	case (2 == 4):
		fmt.Println("this is false")
	case (3 == 3):
		fmt.Println("this is true")
		fallthrough
	case (4 == 3):
		fmt.Println("this is 4==3")
	case (5 == 7):
		fmt.Println("5==7")
	}
	str := "bird"
	switch str {
	case "cat":
		fmt.Println("cat")
	case "dog":
		fmt.Println("dog")
	default:
		fmt.Println("unknown animal")
	}
}
