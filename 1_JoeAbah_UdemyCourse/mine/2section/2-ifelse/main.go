package main

import "fmt"

func main(){
	// temp := 25
	// if temp > 30 {
	// 	fmt.Println("greater")
	// }else{
	// 	fmt.Println("lesser")
	// } 

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	// if accesGranted, ok := userAccess["jane"]; ok && accesGranted{
	// 	fmt.Println("Jane has been granted acces")
	// }
	 accesGranted, ok := userAccess["jane"]
	
	if ok && accesGranted{
		fmt.Println("Jane has been granted access")
	}
}
