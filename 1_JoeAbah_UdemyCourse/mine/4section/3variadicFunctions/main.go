package main

import "fmt"
func sum(number ...int) int {
	//the ...int automatically teslls it its a bunch of string 
	total := 0
	for _, number := range number {
		total += number
	}
	return total
}

// the above is better more put together than having 

func sum1(a int, b int) int {
	return a + b 
}

func main(){
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum1(1, 2))
}