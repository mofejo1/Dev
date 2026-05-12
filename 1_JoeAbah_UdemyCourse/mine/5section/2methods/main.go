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

func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

// if we wanted to create a new functions that doesnt vhnage the initial this is ok, but if we wated to make a chage to the initail struct which will then affect then affct jane we will nee to make use of pointers
func (e *Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

// you will see when i run this it wont reflect in employee because only a copy of it was changed nt he original emplyee
// method (attached to Employee

func (e *Employee) Deactivate() {
	e.IsActive = false
}

// standalone function takes employee as parameter
func deactivate(e *Employee) {
	e.IsActive = true
}

func (e *Employee) changeTime(t time.Time) {
	e.JoinedAt = t
}
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
	fmt.Println(jane.fullName())
	jane.Deactivate()
	fmt.Printf("%+v\n", jane)
	deactivate(&jane)
	fmt.Printf("%+v\n", jane)
	jane.changeTime(time.Now())
}
