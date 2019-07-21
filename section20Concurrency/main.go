package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	goChannel()

	raceCondition()
	raceConditionSolvedByMutex()

	raceConditionSolvedByAtomic()
}

func foo() {
	for i := 0; i < 4; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 4; i++ {
		fmt.Println("bar:", i)
	}
}

func doSomething(x int) int {
	return x * 5
}

func goChannel() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}

func raceCondition() {
	counter := 0

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() //yield the thread
			v++
			counter = v
			wg.Done()
			//fmt.Println("GoRoutines\t", runtime.NumGoroutine())
		}()
	}
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func raceConditionSolvedByMutex() {
	counter := 0

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	const gs = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() //yield the thread
			v++
			counter = v

			wg.Done()
			mu.Unlock()

		}()
		fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	fmt.Println("Counter:", counter)
}

func raceConditionSolvedByAtomic() {
	var counter int64 = 0

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			time.Sleep(time.Second * 2)
			//counter++
			atomic.AddInt64(&counter, 1)

			runtime.Gosched() //yield the thread
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
			//fmt.Println("GoRoutines\t", runtime.NumGoroutine())
		}()
	}
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("raceConditionSolvedByAtomic Counter:", counter)
}
