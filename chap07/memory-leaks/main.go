package main

import (
	"fmt"
	"time"
)

func leakly() {
	ch := make(chan int)

	// never return and never be deallocated
	go func() {
		s := <-ch
		fmt.Println(s)
	}()
}

func timely() {
	timer := time.NewTimer(time.Second * 5)
	// contains a memory leak because never be closed
	ticker := time.NewTicker(time.Second * 1)
	//defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("1s")
			case <-done:
				fmt.Println("done")
				return
			}
		}
	}()
	<-timer.C
	fmt.Println("5s")
	close(done)

}
