package main

import "fmt"

// 1. The interface - defines the CONTRACT
type Person interface {
	GetName() string
}

// 2. The struct - THIS WAS MISSING IN YOUR CODE
type Employee struct {
	ID   int
	Name string
}

// 3. Method on Employee - satisfies the Person interface
func (e Employee) GetName() string {
	return e.Name
}

// 4. Stringer - controls how Employee prints
func (e Employee) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", e.ID, e.Name)
}

func main() {
	joe := Employee{ID: 1, Name: "Joe"}
	fmt.Println(joe) // uses String()
	fmt.Println(joe.GetName())
}
