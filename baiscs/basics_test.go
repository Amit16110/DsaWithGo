package baiscs

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var a = Human{"Smit"}

	val := a.greet()

	fmt.Println("val", val)

	isPerson(&a)
}

func TestInterfaceTypeSecond(t *testing.T) {
	pemp1 := permanent{
		empId:    1,
		basicpay: 5000,
		pf:       20,
	}

	cemp1 := contract{
		empId:    3,
		basicpay: 3000,
	}

	employ := []Salary{&pemp1, &cemp1}

	TotalExpenses(employ)
}
