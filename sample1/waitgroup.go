package main

import (
	"fmt"
	"sync"
	"time"
)

func wgsample() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
		fmt.Println("1st goroutine wake up")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
		fmt.Println("2nd goroutine wake up")
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
}
