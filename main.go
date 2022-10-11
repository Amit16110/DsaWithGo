package main

import (
	"fmt"

	"github.com/amit16110/DsawithGo/striing"
)

// func sliceToChannel(nums []int) <- chan int {
// 	out := make(chan int)  // Unbuffered Channel

// 	go func(){
// 		for _, n := range nums {
// 			out <- n
// 		}

// 		close(out)  // close the channel
// 	}()
// 	return out

// }

// func sq (in <- chan int) <- chan int { // <- chan(read only channel)
// 	out := make()
// }

func main() {
	s1 := "ABACD"
	s2 := "CDABA"

	ans := striing.StringRotationCheck(s1, s2)
	fmt.Println("ans", ans)
}
