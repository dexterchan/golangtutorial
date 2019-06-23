package main

import "fmt"

func main() {
	ex1()
	ex3()
	ex4()
	ex5()
	ex6()
}

const (
	a     = 42
	b int = 42
)

func ex1() {
	x := 100
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 > 43)
	f := (42 < 43)
	fmt.Println(a, b, c, d, e, f)
}
func ex3() {
	fmt.Println(a, b)
}
func ex4() {
	x := 100
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)
	x = x << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", x, x, x)
}
func ex5() {
	x := `it is string literal
			can you see?`
	fmt.Println(x)
}

const (
	y1 = 2017 + iota
	y2 = 2017 + iota
	y3 = 2017 + iota
	y4 = 2017 + iota
)

func ex6() {
	fmt.Println(y1, y2, y3, y4)
}
