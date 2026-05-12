// zero value of a map is nill

package main

import "fmt"

func main() {
	studentScore := map[string]int{
		"Mofe": 99,
		"Toke": 87,
		"Milo": 78,
	}
	fmt.Println(studentScore)
	studentScore["Mofe"] = 100
	fmt.Printf("%+v\n", studentScore)

	//ok chekcks or reflect if mofe exist, and value is stored in mofe
	Mofe, ok := studentScore["Mofe"]; if ok{fmt.Printf("Mofe %+v\n", Mofe)}

	//we can also delete from map 
	delete(studentScore, "Toke")
	fmt.Println(studentScore)

	configs := make(map[string]int)
	if configs != nil{
		fmt.Printf("config is ni")
	}
}
