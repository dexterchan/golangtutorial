package main

import "fmt"

func blockChannel() {
	c := make(chan int, 0) //block channel
	c <- 100
	fmt.Println(<-c)
}
func unblockByThread() {
	c := make(chan int, 0) //block channel
	go func() {
		c <- 100 //push into channel
	}()
	fmt.Println(<-c) //pull from channel
}

func unblockByBufferedChannel() {
	c := make(chan int, 1) //Buffer channel
	c <- 100

	fmt.Println(<-c)
}

func blockByBufferedChannelNotEnoughBuffer() {
	c := make(chan int, 1) //Buffer channel
	c <- 100
	c <- 102
	fmt.Println(<-c)
}
func unblockByBufferedChannelEnoughBuffer() {
	c := make(chan int, 2) //Buffer channel
	c <- 100
	c <- 102
	fmt.Println(<-c) //pull from channel
	fmt.Println(<-c)
}
func main() {
	unblockByBufferedChannelEnoughBuffer()
}
