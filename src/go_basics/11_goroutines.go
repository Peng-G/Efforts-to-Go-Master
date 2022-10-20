package go_basics

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

// use WaitGroup to synchronize goroutines, so that the main thread will wait for the goroutines to finish executing. In this way, we don't need to Sleep() anymore.
var wg = sync.WaitGroup{}
var counter = 0
// Another way to handle racing condition is to use mutex.
var m = sync.RWMutex{} // read-write mutex, only one goroutine can write to the variable, but multiple goroutines can read from the variable.

func Goroutines() {
	fmt.Println("Goroutines")

	sayHello() // this is a normal function call

	go sayHello() // the printed message will be displayed after the main function has finished executing.
	// this is a goroutine call, it will be executed in a separate thread. The main thread will continue to execute the next line of code.
	// The main thread will not wait for the goroutine to finish executing.
	// We will use this higher level of concurrency to execute multiple tasks at the same time.
	time.Sleep(100 * time.Millisecond) // this will make the main thread sleep for 100 milliseconds, so that the goroutine can finish executing.

	// closures, anonymous functions, and goroutines
	var msg = "Hello Go!"
	go func() {
		fmt.Println(msg)
	}()
	// time.Sleep(100 * time.Millisecond)
	// race condition: if we change the value of msg, the goroutine will print the new value, not the old value. Because we are changing the value of msg in the main thread, and the goroutine is reading the value of msg in a separate thread. Untill Sleep() is called, the separate thread will not have time to read the value of msg.
	msg = "Hello Go again!"
	time.Sleep(100 * time.Millisecond)

	// A way to decouple the goroutine from the main thread is to pass in the parameters when calling the function.
	var msg2 = "Hello Go!"
	go func(msg string) {
		fmt.Println(msg)
	}(msg2)
	msg2 = "Hello Go again!"
	time.Sleep(100 * time.Millisecond)

	// use WaitGroup to synchronize goroutines.
	msg3 := "Hello Waitgroup!"
	wg.Add(1) // add 1 to the WaitGroup counter
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // decrement the WaitGroup counter by 1
	}(msg3)
	wg.Wait() // wait for the WaitGroup counter to become 0
	
	// racing condition for WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello2()
		go increment()
	}
	wg.Wait()

	// racing condition for mutex, get locks in each goroutine, so that only one goroutine can write to the variable at a time. But this will cause problems.
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello3()
		go increment2()
	}
	wg.Wait()

	// get locks in the main thread, so that only one goroutine can write to the variable at a time. But this will cause efficiency problems.
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello4()
		m.Lock()
		go increment3()
	}
	wg.Wait()

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // get the number of threads, equal to the number of cores.
	runtime.GOMAXPROCS(16) // set the number of threads to 100. The best number of threads is equal to the number of cores, or sometimes 2 times the number of cores.
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// Practice tips: use "go run -race main.go" to check for racing conditions.
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHello2() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func sayHello3() {
	m.RLock() // read lock
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // read unlock
	wg.Done()
}

func increment2() {
	m.Lock() // write lock
	counter++
	m.Unlock() // write unlock
	wg.Done()
}

func sayHello4() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // read unlock
	wg.Done()
}

func increment3() {
	counter++
	m.Unlock() // write unlock
	wg.Done()
}