package go_basics

import (
	"fmt"
	"sync"
	"time"

)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

const logInfo = "INFO"

var wg0 = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel

func Channels() {
	fmt.Println("Channels")
	// unbuffered channel, only accepts integers, and one message at a time.
	ch := make(chan int) // ch is a bidirectional channel.
	wg0.Add(2)
	// receiving data from channel
	go func(ch <-chan int) {  // <-chan means that this channel is only for receiving data
		i := <-ch  // <- is the operator to receive data from channel
		fmt.Println(i)
		wg0.Done()
	}(ch)
	// sending data to channel
	go func(ch chan<- int) {  // chan<- means that this channel is only for sending data
		ch <- 1  // this step will block until the other goroutine is ready to receive the data. Or there is place int the channel to send data.
		wg0.Done()
	}(ch)
	wg0.Wait()

	// buffered channel, accepts 50 integers.
	ch1 := make(chan int, 50)
	wg0.Add(2)
	// receiving data from channel
	go func() {
		for i := range ch1 {
			fmt.Println(i)
		}
		wg0.Done()
	}()
	// sending data to channel
	go func() {
		ch1 <- 2
		ch1 <- 3
		close(ch1)  // close the channel, no more data can be sent to the channel.
		wg0.Done()
	}()
	wg0.Wait()

	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo,"App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	
	// select statement is used to wait on multiple channel operations.
	go selectLogger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo,"App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}

}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

func selectLogger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}