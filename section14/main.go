package main

import "fmt"

//func (receiver) identifier(parameters) (returns) { code }
func main() {
	bar("pp")
	s := woo("cat")
	fmt.Println(s)
	x, y := mouse("kk", "bb")
	if !y {
		fmt.Println(x)
	}

	variadic(1, 2, 3, 4, 5, 6)
}

//everything in GO is pass by value
func bar(s string) {
	fmt.Printf("Hello, %s\n", s)
}

func woo(s string) string {
	return fmt.Sprintf("Hello from woo,%s", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := (fn == ln)
	b := fmt.Sprintf("%s %s says ''HELLO", fn, ln)
	return b, a
}

//variadic parameter
func variadic(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, v, sum)
	}
}
