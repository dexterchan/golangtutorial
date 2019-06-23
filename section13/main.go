package main

import "fmt"

type Person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	structfunc()
	embeddedfunc()
	annoymousStruct()
}
func structfunc() {
	p := Person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"apple", "mango"},
	}
	fmt.Println(p.first, p.last)

	p2 := Person{
		first:      "Sailor",
		last:       "Moon",
		favFlavors: []string{"water melon", "orange"},
	}

	m := map[string]Person{}
	m[p.last] = p
	m[p2.last] = p2

	for k, v := range m {
		fmt.Println("Key is:", k)
		fmt.Println(v.first, v.last)

		for _, v := range v.favFlavors {
			fmt.Println(v)
		}
	}

}

type Vehicle struct {
	doors int
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}
type Sedan struct {
	Vehicle
	luxury bool
}

func embeddedfunc() {
	t := Truck{
		Vehicle: Vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	s := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "silver",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}

func annoymousStruct() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "pig",
		friends: map[string]int{
			"J": 007,
			"Q": 777,
			"M": 888,
		},
		favDrinks: []string{"spring water", "beer"},
	}

	fmt.Println(s)
	for k, v := range s.friends {
		fmt.Printf("Key:%[1]s value:%[2]v\n", k, v)
	}
	for _, d := range s.favDrinks {
		fmt.Println(d)
	}
}
