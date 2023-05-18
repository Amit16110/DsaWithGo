package array

import "fmt"

func myfunc(arr []int) {
	for i := 0; i < 7; i++ {
		arr = append(arr, i)
	}

}

func callingMyFunc() {

	slicArr := []int{}
	// for i := 0; i < 21; i++ {
	// 	slicArr = append(slicArr, i)
	// }
	// fmt.Println(slicArr)
	myfunc(slicArr)
	fmt.Println("let's see", slicArr)
}
