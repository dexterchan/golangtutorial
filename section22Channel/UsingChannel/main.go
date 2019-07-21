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
	c <- 100
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
