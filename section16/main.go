package main

import "fmt"

func main() {
	a := 99
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Printf("Reference:%[1]v Dereference:%[2]v\n", b, *b)

	*b = 100
	fmt.Printf("Reference:%[1]v Dereference:%[2]v\n", b, *b)

	mutateValue(b)
	fmt.Printf("Reference:%[1]v Dereference:%[2]v\n", b, *b)

}

func mutateValue(v *int) {
	*v = 200
}
