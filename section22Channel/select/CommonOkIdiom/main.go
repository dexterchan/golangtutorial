package main

import (
	"fmt"
)

func withoutOk() {
	c := make(chan int)

	go func() {
		c <- 100
	}()
	v := <-c
	fmt.Println(v)

	close(c)

	v = <-c
	fmt.Println(v)
}

func withOk() {
	c := make(chan int)

	go func() {
		c <- 100
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func main() {
	withoutOk()
	withOk()
}
