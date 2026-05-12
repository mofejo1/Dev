package main

import "fmt"

// CONTRACT - any type that has GetName() is considered a Person
type Person interface {
	GetName() string
}

// BLUEPRINTS - the data shapes
type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

// SATISFYING THE CONTRACT - attaching GetName() to each struct via receiver
// Employee now fulfills the Person contract
func (e Employee) GetName() string {
	return e.Name
}

// BusinessPerson now fulfills the Person contract
func (b BusinessPerson) GetName() string {
	return b.Name
}

// STRINGER - controls how each type prints via fmt.Println
// defined once on the type, every Employee variable uses this format automatically
func (e Employee) String() string {
	return fmt.Sprintf("Employee[ID:%d, Name:%s]", e.ID, e.Name)
}

func (b BusinessPerson) String() string {
	return fmt.Sprintf("BusinessPerson[ID:%d, Name:%s]", b.ID, b.Name)
}

// INTERFACE FUNCTIONS - accept any type that satisfies Person
// one function works for Employee, BusinessPerson, or any future type
func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func greetPerson(p Person) {
	fmt.Println("Hello,", p.GetName())
}

func main() {
	joe := Employee{ID: 1, Name: "Joe"}
	jane := BusinessPerson{ID: 2, Name: "Jane"}

	// direct method call - always available once contract is satisfied
	fmt.Println(joe.GetName())
	fmt.Println(jane.GetName())

	// fmt.Println uses String() automatically - no extra work needed
	fmt.Println(joe)
	fmt.Println(jane)

	// interface functions - one function handles both types
	displayPerson(joe)
	displayPerson(jane)

	greetPerson(joe)
	greetPerson(jane)
}
