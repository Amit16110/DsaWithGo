package baiscs

import "fmt"

// what is the interface?

// decalared the interface
type Person interface {
	greet() string
}

type Human struct {
	Name string
}

func (h *Human) greet() string {
	return "Hi, I'm" + h.Name
}

func isPerson(h Person) {
	fmt.Println(h.greet())
}

// TYPE 2

type Salary interface {
	calculateSalary() int
}

type (
	permanent struct {
		empId    int
		basicpay int
		pf       int
	}
	contract struct {
		empId    int
		basicpay int
	}
	freelenance struct {
		empId       int
		ratePerHour int
		totalHours  int
	}
)

func (p *permanent) calculateSalary() int {
	return p.basicpay + p.pf
}
func (c *contract) calculateSalary() int {
	return c.basicpay
}
func (f *freelenance) calculateSalary() int {
	return f.totalHours * f.ratePerHour
}

func TotalExpenses(s []Salary) {
	expense := 0
	for _, v := range s {
		fmt.Println("value insdeide", v.calculateSalary())
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}
