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

	var unfold []int = []int{1, 2, 3}
	variadic(1, 2, 3, 4, 5, 6)
	variadic(unfold...)
	playdefer()

	setupSecretAgent()

	anonymousfunc()

	sqFunc := returnFunc()
	fmt.Printf("%T sqFunc:%d\n", sqFunc, sqFunc(10))

	demoCallback()

	runScope()
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

func foo() {
	fmt.Println("foo")
}
func bar2() {
	fmt.Println("bar")
}
func playdefer() {
	defer foo()
	bar2()
	fmt.Println("closing playdefer")
}

type Pet struct {
	name string
}
type Person struct {
	first string
	last  string
}

type SecretAgent struct {
	Person
	ltk bool
}

//attach a method to a type struct
func (s SecretAgent) speak() {
	fmt.Println("I am ", s.last, ",", s.first, s.last)
}

func (p Person) speak() {
	fmt.Println("Hello, I am ", p.first, p.last)
}

func setupSecretAgent() {
	sa1 := SecretAgent{
		Person: Person{
			"James",
			"Bond",
		},
		ltk: false,
	}
	pp := Person{
		"Bob",
		"X",
	}
	fmt.Println(sa1)
	sa1.speak()

	/*pet1 :=
	Pet{
		"dog",
	}*/
	h := HumanInterface(sa1)
	fmt.Println(h)
	interface_polymorphism(sa1)
	interface_polymorphism(pp)
	//interface_polymorphism(pet1)
}

type HumanInterface interface {
	speak()
}

func interface_polymorphism(h HumanInterface) {

	switch h.(type) {
	case Person:
		fmt.Println("it is person", h.(Person)) //assert it is Person
	case SecretAgent:
		fmt.Println("It is secret agent", h.(SecretAgent))
	default:
		fmt.Println("it is unknown")
	}
}

func anonymousfunc() {
	a := func(x int) {
		fmt.Println("Anonymous func run", x)
	}
	a(100)
}

func returnFunc() func(x int) int {
	return func(x int) int {
		return x * x
	}
}

func demoCallback() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("odd sum:", odd(sumUp, x...))
}
func sumUp(x ...int) int {
	accum := 0
	for _, v := range x {
		accum += v
	}
	return accum
}

func odd(f func(...int) int, x ...int) int {
	var temp []int
	for _, v := range x {
		if v%2 == 1 {
			temp = append(temp, v)
		}
	}
	return f(temp...)
}

func incrementor() func() int {
	var x = 10

	return func() int {
		x++
		return x
	}

}

func runScope() {
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
