package main

import "fmt"

func sendOnlyChannel() {
	c := make(chan<- int, 2) //send only channel

	c <- 42
	c <- 43

	//fmt.Println(<-c)
	//fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("send-only channel%T\n", c)
}

func receiveOnlyChannel() {
	c := make(<-chan int, 2) //receive only channel

	//c <- 42
	//c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("send-only channel%T\n", c)
}

func main() {
	sendOnlyChannel()
	receiveOnlyChannel()
}
