//Panic and recovery

package main

import "fmt"
func mightPanic(shouldPanic bool){
	if shouldPanic {
		panic("something went wrong in mightPanic")
	}
	fmt.Println("This function executed without a panic")
}

// if panic , how should we recover 
func recoverable(){
	defer func(){
		//recover() is a built in function in go 
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	mightPanic(false)
}

func main(){
 	recoverable()
	// you can panic the programme manually 
	//panic("somthing bad happened")
}