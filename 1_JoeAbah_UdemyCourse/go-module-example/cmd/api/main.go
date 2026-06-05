package main

import (
	"example/app/models"
	"fmt" // inporting internal packages
)

func main() {
	var jane models.User
	jane.Name = "Jane"
	fmt.Println("Hello, World!", jane)
}
