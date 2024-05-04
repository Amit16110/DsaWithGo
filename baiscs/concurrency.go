package baiscs

import (
	"fmt"
	"sync"
)

func print(i int) int {
	return i
}

// pass the channel into the function argu

func printArg(i int, ch chan int) {
	ch <- i
}

func PrintWithWait() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Println("value", print(i))
		}(i)

	}
	wg.Wait()
}
func printWithChannel() {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- print(i)
		}(i)
	}

	for v := range c {
		fmt.Println("print ==>", v)
	}
	close(c)

}

// by using this model we get the error.
func printWithChannelWithArg() {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go printArg(i, c)
	}
	// defer close(c)

	for v := range c {
		fmt.Println("print ==>", v)
	}

}

func printWithChanWaitGroup(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- i
}

func betterSolution() {
	c := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go printWithChanWaitGroup(i, c, &wg)

		go func() {
			wg.Wait()
			close(c)
		}()
	}

	for v := range c {
		fmt.Println("print ==>", v)
	}

}

// Practise concurency with the wait groups

// sync in between two functions it print hello than world.
func hello(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	fmt.Println("hello")

	ch <- true
}
func world(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()

	<-ch
	fmt.Println("world")
}

func syncronizeInBetweenTwoGo() {

	var wg sync.WaitGroup
	ch := make(chan bool)

	wg.Add(2)

	go hello(&wg, ch)
	go world(&wg, ch)

	wg.Wait()

}

func changeOneVariableInSchronization() {
	// we have a varibale x which we need to change by the go routine
	x := 0

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	go func(ch chan int) {
		defer wg.Done()
		ch <- 5
	}(ch)

	go func(ch chan int) {
		defer wg.Done()
		ch <- 10
	}(ch)

	// go func(ch chan int) {
	// 	x = <-ch
	// }(ch)
	x = <-ch

	wg.Wait()

	fmt.Println("x", x)
}

func producer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// close(ch)
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func callConPro() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go producer(&wg, ch)
	go consumer(&wg, ch)

	wg.Wait()

}
