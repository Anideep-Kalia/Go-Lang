package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang-  LearnCodeOnline.in")

	myCh := make(chan int, 2) // so we are defining that there are 2 channel or 2 space in stack of channel anything more than that value would be erased and filled again
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// R ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) { // so the arrow plays crucial role in defing role of function for the channel as we can see here if there was no arrow then function can read and write both

		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		// close() : is a send only function for channel so you can't write it inside Read only function
		close(myCh) // so if we now try to get value we will get 0 and not nill
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
