package regularexpression
package main 

import "regexp"

func main() {
	text1 := "Hellow world! Welcome to Go!"
	regGo, err := regexp.MustCompile(`Go`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		od.Exit(1)
	}
	fmt.Printf("Text '%s' contains 'Go': %t\n", text1, regGo.MatchString(text1))

}