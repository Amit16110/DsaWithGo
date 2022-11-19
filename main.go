package main

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

func sort(asse , des []int) []int {
	asselen := len(asse)
	deslen := len(des)
	ans := make([]int,asselen)
	if asselen && deslen < 1 {
		return []int
	}

	for i:= 0, i < asselen; i++{
		for j := 0; j < deslen; j++ {
			if asse[i] < des[j]{
				ans = append(ans, )
			}
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []int{10, 9, 8, 7, 3}


}
