package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

//Model subclass by embedding
type SecretAgent struct {
	Person
	licenseToKill bool
}

func main() {
	namedStruct()
	annoymousStruct()
}
func namedStruct() {
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   35,
	}
	p2 := Person{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	sa1 := SecretAgent{
		Person:        p1,
		licenseToKill: true,
	}
	//can skip "Person" when field is unique is embedded struct
	fmt.Println(sa1.Person.first, sa1.last, sa1.licenseToKill)
}

func annoymousStruct() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)
}
