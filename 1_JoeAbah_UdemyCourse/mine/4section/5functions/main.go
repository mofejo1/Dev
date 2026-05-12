package main

import (
	"errors"
	"fmt"
	"strings"
)
func devide(a, b int)(int, error){
	if b == 0 {
		return 0, errors.New("divide by zero ")
	}
	return a/b, nil
}


func splitName(fullName string)(firstName string, lastName string){
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]
	return 
}


func main(){
	value, err := devide(10, 0)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(value)
	}
	firstName, lastName := splitName("Joseph Abah")
	fmt.Println(firstName, lastName)
}