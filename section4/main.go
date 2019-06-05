package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello world")
	calculate()
	fmt.Println("Number of bytes:", n)
	fmt.Println(err)
	declareAndAssign()
	exploreType()
	fmtShow()
	typestudy()
}

func calculate() {
	var c float64 = 97.0
	c = c
	num := 1000
	accum := 0.0
	for i := 0; i < num; i++ {
		if i%2 == 0 {
			accum += 1
		}

	}
	n, _ := fmt.Println(accum)
	fmt.Println("Number of bytes:", n)
}

//Declare the variable outside a function
var global_x = 100

//Declare a variable with the identifier floating
// and the type is float64
// and assign ZERO to it
var floating float64

func declareAndAssign() {
	//Declare a value and assign a value at the same time
	x := 100
	fmt.Println("initalize x:", x)
	x = 200
	fmt.Println("Assign new value of x:", x)

	fmt.Println("Value of floating", floating)
}

func exploreType() {
	fmt.Println("type usage")
	var y float32 = 42
	var str string = "Hello World"
	var str2 string = "pp said, 'apple'"
	fmt.Println("Explore type")
	fmt.Printf("%T\n", y)
	fmt.Printf("%T \n", str)
	fmt.Println(str2)
}

func fmtShow() {
	fmt.Println("fmt")
	var y int = 900
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#[1]x\n", y)

	s := fmt.Sprintf("%#[1]x\n", y)
	fmt.Println(s)
}

func typestudy() {
	fmt.Println("Custom type")
	type dog int
	var a int = 100
	var b dog = 101
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b) //Conversion
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
