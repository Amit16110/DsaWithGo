package main

import (
	"fmt"
)

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

func moveNegative(arr []int) {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[right] < 0 {
			right -= 1
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			left += 1
		}
	}
	fmt.Println("data===>", arr)
}

func main() {

}
