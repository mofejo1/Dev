package main

import "fmt"

//composition is a design principle in which a class is composed of one or more objects of other classes, rather than inheriting from them.
// This allows for greater flexibility and modularity in the design of the class,
// as it can use the functionality of the composed objects without being tightly coupled to
//  their implementation. In Go, composition is often achieved through the use of struct embedding,
// where one struct can embed another struct to gain access to its fields and methods.

type Address struct {
	street string
	city   string
	state  string
	zip    string
}

func (a Address) FullAddress() string {
	if a.street == "" && a.city == "" {
		return "No adress provided"
	}
	return fmt.Sprintf("%s, %s, %s %s", a.street, a.city, a.state, a.zip)
}

type Customer struct {
	CustomerID     int
	Name           string
	Email          string
	BillingAdress  Address //embedded
	shippingAdress Address // embedded
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("Billing Address: %s\n", c.BillingAdress.FullAddress())
	fmt.Printf("Shipping Address: %s\n", c.shippingAdress.FullAddress())
}

func main() {
	cust1 := Customer{
		CustomerID: 1001,
		Name:       "Gadget Corp",
		Email:      "sales@gads1@gadgetcorp.com",
		BillingAdress: Address{
			street: "123 Tech Park",
			city:   "Innovateville",
			state:  "CA",
			zip:    "94025",
		},
		shippingAdress: Address{
			street: "456 Innovation Blvd",
			city:   "Silicon Valley",
			state:  "CA",
			zip:    "94025",
		},
	}
	cust1.PrintDetails()

	fmt.Println("------- customer with the same billing and shipping add")
	mainAdress := Address{
		street: "789 Main St",
		city:   "Tech City",
		state:  "CA",
		zip:    "94025",
	}
	cust2 := Customer{
		CustomerID:     1002,
		Name:           "Tech Solutions",
		Email:          "john.doe@email.com",
		BillingAdress:  mainAdress,
		shippingAdress: mainAdress,
	}
	cust2.PrintDetails()

}

//inheritance vs composition
//Inheritance is a fundamental object-oriented programming principle where a new class (called a subclass or child class) is created based on an existing class (called a superclass or parent class). The subclass inherits the properties and behaviors of the superclass, allowing for code reuse and the creation of a hierarchical relationship between classes. Inheritance promotes a "is-a" relationship, where the subclass is a specialized version of the superclass.

//Composition, on the other hand, is a design principle where a class is composed
