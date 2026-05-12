package main

import "fmt"

const(
	sunday = iota + 1
	monday 
	teusday
)

func main(){
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(teusday)
}