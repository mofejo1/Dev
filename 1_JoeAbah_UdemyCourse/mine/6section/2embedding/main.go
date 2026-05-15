package main

import "fmt"

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

type contactInfo struct {
	email   string
	phone   string
	address Address
}

func (c contactInfo) ContactDetails() string {
	return fmt.Sprintf("Email: %s\nPhone: %s\nAddress: %s", c.email, c.phone, c.address.FullAddress())
}

type company struct {
	name string
	Address
	contactInfo
	BussinessType string //shadows the one in contactInfo
}

func (c company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.name)
	fmt.Printf("loaction: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.street)
	fmt.Printf("Email (promoted): %s\n", c.email)
	fmt.Printf("Business Type: %s\n", c.BussinessType)

}

func main() {
	comp := company{
		name: "Tech Innovators Inc.",
		Address: Address{
			street: "123 Innovation Drive",
			city:   "Tech City",
			state:  "CA",
			zip:    "94025",
		},
		contactInfo: contactInfo{
			email: "contact@inovate.com",
			phone: "555-1234",
		},
		BussinessType: "Software Development",
	}
	comp.GetProfile()

}
