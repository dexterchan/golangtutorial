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

	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl)
	changeSlic(sl)
	fmt.Println(sl)
}

func mutateValue(v *int) {
	*v = 200
}

func changeSlic(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * 2
	}
}
