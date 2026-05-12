package main

import "fmt"

// recursion


func factorial (n int) int {
	// for recursion base state is very important 
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}



// you can return a function as an argument 
func intSeq() func() int {
	i := 0
	return func() int{
		i++
		return 1
	}
}

// we can also write a function into a variable 

func main(){
	fmt.Println(factorial(5))
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	logger := func(message string){
	fmt.Println(message)
}
	logger("Hello world")

}