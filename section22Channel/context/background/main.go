package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func cancelExample() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * time.Duration(200))
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())
}

func cancelShow() {
	ParentCtx := context.Background()

	fmt.Println("context:\t", ParentCtx)
	fmt.Println("context err:\t", ParentCtx.Err())
	fmt.Printf("context type:\t%T\n", ParentCtx)
	fmt.Println("----------")

	ctx1, cancelFunc := context.WithCancel(ParentCtx)
	fmt.Println("context:\t", ctx1)
	fmt.Println("context err:\t", ctx1.Err())
	fmt.Printf("context type:\t%T\n", ctx1)
	fmt.Println("cancel:\t\t", cancelFunc)
	fmt.Printf("cancelFunc type: %[1]T\n", cancelFunc)
	fmt.Println("----------")

	cancelFunc()
	fmt.Println("-------After calling cancel----------")
	fmt.Println("context:\t", ctx1)
	fmt.Println("context err:\t", ctx1.Err())
	fmt.Printf("context type:\t%T\n", ctx1)
	fmt.Println("cancel:\t\t", cancelFunc)
	fmt.Printf("cancelFunc type: %[1]T\n", cancelFunc)
	fmt.Println("----------")
}

func main() {
	cancelShow()
	cancelExample()
}
