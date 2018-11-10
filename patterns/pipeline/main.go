package main

import (
	"fmt"
)

func generator(done <-chan interface{}, integers ...int) <-chan int {
	fmt.Println("[generator] start func")
	intStream := make(chan int, len(integers))

	go func() {
		defer close(intStream)
		for _, i := range integers {
			fmt.Printf("[generator] %v start\n", i)
			select {
			case <-done:
				fmt.Println("[generator] done")
				return
			case intStream <- i:
				fmt.Printf("[generator] %v send to stream\n", i)
			}
			fmt.Println("[generator] end select in goroutine")
		}
		fmt.Println("[generator] end for in goroutine")
	}()

	fmt.Println("[generator] return channel")
	return intStream
}

func multiply(
	done <-chan interface{},
	intStream <-chan int,
	multiplier int,
) <-chan int {
	fmt.Println("[multiply] start func")
	multipliedStream := make(chan int)

	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			fmt.Printf("[multiply] %v start\n", i)
			select {
			case <-done:
				fmt.Println("[multiply] done")
				return
			case multipliedStream <- i * multiplier:
				fmt.Printf("[multiply] %v * %v send to stream\n", i, multiplier)
			}
			fmt.Println("[multiply] end select in goroutine")
		}
		fmt.Println("[multiply] end for in goroutine")
	}()

	fmt.Println("[multiply] return channel")
	return multipliedStream
}

func add(
	done <-chan interface{},
	intStream <-chan int,
	additive int,
) <-chan int {
	fmt.Println("[add] start func")
	addedStream := make(chan int)

	go func() {
		defer close(addedStream)
		for i := range intStream {
			fmt.Printf("[add] %v start\n", i)
			select {
			case <-done:
				fmt.Println("[add] done")
				return
			case addedStream <- i + additive:
				fmt.Printf("[add] %v + %v send to stream\n", i, additive)
			}
			fmt.Println("[add] end select in goroutine")
		}
		fmt.Println("[add] end for in goroutine")
	}()

	fmt.Println("[add] return channel")
	return addedStream
}

func main() {
	fmt.Println("[main] start func")
	done := make(chan interface{})
	defer close(done)

	fmt.Println("[main] define intStream")
	intStream := generator(done, 1, 2, 3, 4)
	fmt.Println("[main] define pipeline")
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	fmt.Println("[main] wait for pipeline...")
	for v := range pipeline {
		fmt.Printf("[main] received %v from pipeline\n", v)
		fmt.Println(v)
	}
	fmt.Println("[main] end func")
}
