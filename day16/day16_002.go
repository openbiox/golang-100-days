//From <https://studygolang.com/articles/12560?fr=sidebar>
package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}
type Contract struct {
	empId    int
	basicpay int
}
type Permanent struct {
	empId    int
	basicpay int
	jj       int //Bonus
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("total $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 3000, 10000}
	pemp2 := Permanent{2, 3000, 20000}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
