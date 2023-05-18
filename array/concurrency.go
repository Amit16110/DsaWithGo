package array

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
