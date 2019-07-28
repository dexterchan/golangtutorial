package main

import (
	"fmt"
	"sync"
)

func gen(step int, concurrent int) <-chan int {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(concurrent)

	go func() {
		for n := 0; n < concurrent; n++ {
			go func() {
				for i := 0; i < step; i++ {
					c <- i
				}
				wg.Done()
			}()
		}
		wg.Wait()
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
	c := gen(10, 10)
	receive(c)
}
