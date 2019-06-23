package main

import "fmt"

const a int = 1
const b float64 = 0.01
const c string = "Hello"

const (
	aa = iota //iota only reset when it hit const keyword
	bb
	cc
	dd = iota
	ee
	ff
)

func main() {
	s := "Hello World"

	bs := []byte(s) //conversion of string to bytes
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%U\n", s[i])
	}
	for i, v := range s {
		fmt.Printf("at index position %[1]d we have hex %#[2]x\n", i, v)
	}

	fmt.Println(aa)
	fmt.Println(dd)

	bitshift()
}

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func bitshift() {

	fmt.Printf("%d\t\t%[2]b\n", kb, kb)
	fmt.Printf("%d\t\t%[2]b\n", mb, mb)
	fmt.Printf("%d\t\t%[2]b\n", gb, gb)
}
