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
	// arr := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}}
	// result := striing.CountMatchOptimse(arr, "type", "phone")
	// fmt.Println("result===>", result)

	// fmt.Println("result =>", striing.Anagram("aacc", "ccac"))
	fmt.Println("LPS=>", striing.LongestCommonSubsequence("abcde", "abef"))
}
