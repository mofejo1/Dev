func main
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

func (c contactInfo) ContactDetails() string{
	return fmt.Sprintf("Email: %s\nPhone: %s\nAddress: %s", c.email, c.phone, c.address.FullAddress())
}