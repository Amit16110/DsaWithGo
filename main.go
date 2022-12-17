package main

import linkedlist "github.com/amit16110/DsawithGo/linkedList"

// "github.com/amit16110/DsawithGo/array"

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
	// fmt.Println("result>>", striing.EditDistance("horse", "ros"))
	linkedlist.LinkedListOperations(10)

}
