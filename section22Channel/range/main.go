package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)
	//receive
	bar(c) //make sure it run

	fmt.Println("Exit")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) //need to close otherwise it block forever
}

func bar(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
