package main

import (
	"fmt"
	"sync"
	"time"
)

// Main function runs in it's own go routine
func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(4)
	// spinning off each of the sleep functions in their own go routiens
	go func() {
		sleep100ms()
		wg.Done()
	}()
	go func() {
		sleep200ms()
		wg.Done()
	}()
	go func() {
		sleep300ms()
		wg.Done()
		wg.Done()
	}()
	go func() {
		sleep400ms()
		// wg.Done()
	}()
	fmt.Println("elapsed time before wait:", time.Since(now))
	wg.Wait()
	fmt.Println("elapsed time after wait:", time.Since(now))
	msgChannel := make(chan string)
	go func() {
		sleep400ms()
		msgChannel <- "done"
	}()
	fmt.Println("waiting for msgChannel")
	<-msgChannel
	fmt.Println("msgChannel joined:", time.Since(now))
}

func sleep100ms() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("100ms")
}

func sleep200ms() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("200ms")
}

func sleep300ms() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("300ms")
}

func sleep400ms() {
	time.Sleep(400 * time.Millisecond)
	fmt.Println("400ms")
}
