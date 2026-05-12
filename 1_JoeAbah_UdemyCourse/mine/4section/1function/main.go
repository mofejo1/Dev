package main

//functions

import "fmt"

func greet(name string){
	fmt.Print("Hello ", name)
}

func add (a int, b int){
	fmt.Println(a, b, a+b)
}

func main(){
	greet("Mofe")

}


// func main(){
// numbers := []int{3, 7, 1, 9, 4}
// var high int

// for _, num := range numbers {
// if num > high {
// 	high = num
// }
// }
// fmt.Println(high)
// }

