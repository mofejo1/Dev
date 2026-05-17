package main

import (
	"fmt"
	"strings"
)

// the strings package coantains various functions for manipulating strings, including Clone, which creates a copy of a string. In this example, we create a string s1 and then use strings.Clone to create a new string s2 that is a copy of s1. Finally, we print both strings to the console.
func main() {
	s1 := "abc"
	s2 := strings.Clone(s1)
	fmt.Println(s1, s2)

	b := strings.Builder{}
	b.WriteString("Hello, ")
	b.Write([]byte("word!"))
	b.WriteString("world!")
	fmt.Println(b.String())

	// c := strings.Reader(strings.NewReader("Hello, world!"))
	// buf := make([]byte, 5)
	// n, err := c.Read(buf)
	// if err != nil {
	// 	fmt.Printf("Error reading from string reader: %v\n", err)
	// 	return
	// }
	// fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
}

// we also have types e.g the strings.Builder type, which is a type that provides a way to efficiently build strings by appending to them. In this example, we create a new strings.Builder and use the WriteString method to append the string "Hello, " to it. Then we append "world!" to the builder. Finally, we call the String method to get the final string and print it to the console.
