package main

import "fmt"

func gen(step int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < step; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
func receive(rec <-chan int) {
	for i := range rec {
		fmt.Println(i)
	}
}
func main() {
	c := gen(10)
	receive(c)
}
