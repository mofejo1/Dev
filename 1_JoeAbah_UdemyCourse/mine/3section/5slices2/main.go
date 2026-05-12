package main

import "fmt"




func main(){
fmt.Println("more on slices ")
original := []int{1, 2, 3, 4, 5, 6, 7, 8}
fmt.Println(original)
s1 := original[2:5] //5 is not included 
s2 := original[:4] //excluding 4
fmt.Println(s1,s2)
}