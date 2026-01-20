package main

import (
	"fmt"
	"sync"
	"time"
)

/*
========================
SECTION 1: BASICS
========================
*/

func basicChannel() {
	// Unbuffered channel
	ch := make(chan int)

	// Sender goroutine
	go func() {
		// This SEND will BLOCK until someone RECEIVES
		ch <- 10
	}()

	// Receiver (main goroutine)
	val := <-ch // blocks until sender sends
	fmt.Println("Basic channel value:", val)
}

/*
Key rule:
- Unbuffered channel = send and receive must meet at the same time
*/

/*
========================
SECTION 2: BUFFERED CHANNEL
========================
*/

func bufferedChannel() {
	// Buffered channel with capacity 2
	ch := make(chan int, 2)

	// These sends DO NOT block (buffer has space)
	ch <- 1
	ch <- 2

	// Third send would BLOCK (buffer full)
	// ch <- 3 // uncomment to see deadlock

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

/*
Key rule:
- Buffered channel = queue
- No overwriting
- Blocks only when buffer is full (send) or empty (receive)
*/

/*
========================
SECTION 3: CHANNEL DIRECTION
========================
*/

func sender(ch chan<- int) {
	// Can ONLY send
	ch <- 100
	// close(ch) // allowed (sender owns channel)
}

func receiver(ch <-chan int) {
	// Can ONLY receive
	val := <-ch
	fmt.Println("Directional receive:", val)

	// ch <- 5        // ❌ compile error
	// close(ch)      // ❌ compile error
}

func channelDirection() {
	ch := make(chan int)

	go sender(ch)
	receiver(ch)
}

/*
Key rule:
- Directional channels are COMPILE-TIME SAFETY
- They prevent accidental misuse
*/

/*
========================
SECTION 4: CLOSING CHANNELS
========================
*/

func closingChannel() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch) // signal: no more values
	}()

	// Receive until channel is closed
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received:", val)
	}
}

/*
Key rules:
- ONLY sender should close
- Closing is a SIGNAL, not cleanup
- Receiving from closed channel:
  - remaining values
  - then zero value + ok=false
*/

/*
========================
SECTION 5: RANGE OVER CHANNEL
========================
*/

func rangeOverChannel() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Automatically stops when channel is closed
	for val := range ch {
		fmt.Println("Range received:", val)
	}
}

/*
Best practice:
- Prefer range when consuming until close
*/

/*
========================
SECTION 6: WAITGROUP + CHANNELS
========================
*/

func waitGroupExample() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		ch <- 42
	}()

	go func() {
		defer wg.Done()
		val := <-ch
		fmt.Println("WaitGroup value:", val)
	}()

	wg.Wait()
}

/*
Channels synchronize DATA
WaitGroups synchronize LIFETIME
They solve DIFFERENT problems
*/

/*
========================
SECTION 7: DEADLOCK EXAMPLE
========================
*/

func deadlockExample() {
	ch := make(chan int)

	// ❌ Deadlock: no receiver
	// ch <- 1

	_ = ch
}

/*
Rule:
- Every send must have a receiver
- Every receive must have a sender
*/

/*
========================
SECTION 8: SELECT STATEMENT
========================
*/

func selectExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	select {
	case v := <-ch1:
		fmt.Println("Received from ch1:", v)
	case v := <-ch2:
		fmt.Println("Received from ch2:", v)
	}
}

/*
select:
- waits on multiple channels
- picks the first ready one
- essential for timeouts & fan-in
*/

/*
========================
SECTION 9: CHANNEL AS SIGNAL (NO DATA)
========================
*/

func signalChannel() {
	done := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		close(done) // signal completion
	}()

	<-done
	fmt.Println("Signal received, work done")
}

/*
Best practice:
- Use chan struct{} when no data is needed
*/

/*
========================
MAIN
========================
*/

func main() {
	basicChannel()
	bufferedChannel()
	channelDirection()
	closingChannel()
	rangeOverChannel()
	waitGroupExample()
	selectExample()
	signalChannel()
}
