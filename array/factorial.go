package array

import "fmt"

func Factorial(numb int) {
	//
	fact := 1

	for numb != 0 {
		fact = fact * numb
		numb--
	}
	fmt.Println("factorial", fact)
}
