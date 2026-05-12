package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

// when you have a struct you can create a value typoe , eg. jane that is atype of employee
// func main (){

// }

// & — gives you the address (where it lives in memory)
// * — gives you the value at that address (what's inside)
func main() {
	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "Night",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
	fmt.Println(jane)
	fmt.Println(jane.ID)
	fmt.Println(jane.LastName)
	fmt.Println(jane.Position)
	fmt.Println(jane.Salary)

	Mofe := NewEmployee(2, "Jimoh", "Onisemo", "Day", 2000, true, time.Now())

	fmt.Println(Mofe.LastName)
	fmt.Println(Mofe)
	fmt.Println(Mofe.FirstName)

	Mofe.Salary = 50000
	mofePtr := &Mofe
	// this brings out the salry stored in the memory address of Mofe
	fmt.Println(mofePtr.Salary)
	fmt.Println(mofePtr)
	mofePtr.IsActive = true
	// basically we can use strored the address of Mofe in mofePtr and use it to change the value of Mofe
	fmt.Println(Mofe)
}

func NewEmployee(id int, firstName, lastName, position string, salary int, isActive bool, joinedAt time.Time) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  joinedAt,
	}
}
