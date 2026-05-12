package main

import (
	"fmt"
	"os"
)

// defer


func main(){
	//another example is writing file in 
	file, err := os.Create("./defer.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()




	defer func(){
		fmt.Println("before the return of main")
	}()// the () is calling it immediately after declaring it 
	simpleDefer()
	defer func(){
		fmt.Println("after the return of main")
	}()
}

func simpleDefer(){
	fmt.Println("function simpleDefer: start")
	defer fmt.Println("function simpleDefer: defered1")
	fmt.Println("function simpleDefer: middle")
	defer fmt.Println("function simpleDefer: defered")
}