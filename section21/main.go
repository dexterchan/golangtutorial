package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func exercise1() {

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	go func() {
		fmt.Println("Hello from A")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from B")
		wg.Done()
	}()
	fmt.Println("exit now")
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end CPU\t", runtime.NumCPU())
	fmt.Println("end GoRoutines\t", runtime.NumGoroutine())
}

type Person struct {
	name string
}

func (p *Person) speak() {
	fmt.Printf("My name is %[1]v\n", p.name)
}

type Human interface {
	speak()
}

func saySomething(p *Person) {
	p.speak()
}
func exercise2() {
	p := Person{
		name: "Peter Parker",
	}
	//saySomething(p) //not working when passing a value
	saySomething(&p)
}

func racecondition_withoutWaitGroup() {
	inc := 0
	numOfRoutine := 1000

	var wg sync.WaitGroup
	wg.Add(numOfRoutine)
	for i := 0; i < numOfRoutine; i++ {
		go func() {
			inc++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final number is:%d\n", inc)
}

func exercise3() {
	racecondition_withoutWaitGroup()
}

func solevRacecondition_withMutex() {
	inc := 0
	numOfRoutine := 1000

	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(numOfRoutine)
	for i := 0; i < numOfRoutine; i++ {
		go func() {
			mu.Lock()
			inc++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Final number is:%d\n", inc)
}

func exercise4() {
	solevRacecondition_withMutex()
}

func solevRacecondition_withAtomic() {
	var inc int64 = 0
	numOfRoutine := 1000

	var wg sync.WaitGroup

	wg.Add(numOfRoutine)
	for i := 0; i < numOfRoutine; i++ {
		go func() {
			atomic.AddInt64(&inc, 1)
			fmt.Println(atomic.LoadInt64(&inc))
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Printf("Final number is:%d\n", inc)
}

func exercise5() {
	solevRacecondition_withAtomic()
}

type Counter struct {
	value int64
}
type sumInterface interface {
	increment()
}

func (c *Counter) increment() {
	c.value++
}

func solveRacecondition_withAtomic2() {

	var incS = Counter{
		value: 0,
	}

	numOfRoutine := (1000)
	//var s1 = sumInterface(&incS)
	//cc := s1.(*Counter)
	//cc.increment()
	var wg sync.WaitGroup

	//fmt.Printf("before: %T\n", inci)
	//fmt.Printf("after: %T\n", atomCounter.Load().(*Counter))
	var atomCounter atomic.Value
	atomCounter.Store(&incS)
	wg.Add(numOfRoutine)
	var mu sync.Mutex

	//It feature a more frequent read, but a lower frequency of write
	for i := 0; i < numOfRoutine; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			var s = atomCounter.Load().(*Counter)
			s.increment()
			time.Sleep(2 * time.Millisecond)
			atomCounter.Store(s)
			wg.Done()

		}()
	}

	go func() {
		var s *Counter
		s = atomCounter.Load().(*Counter)
		for s.value < int64(numOfRoutine) {
			fmt.Println("Now:", s.value)
			s = atomCounter.Load().(*Counter)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("Now:", s.value)
	}()

	wg.Wait()
	finalValue := (atomCounter.Load().(*Counter))
	fmt.Printf("Final number is:%d\n", finalValue.value)
}
func exerciseX() {
	solveRacecondition_withAtomic2()
}
func main() {

	exerciseX()

	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
