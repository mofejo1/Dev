package main

import (
	"example/internal/color"
	"fmt"
)

func main() {
	redText := "this is a red Text"
	greenText := "this is a green Text"
	blueText := "this is a blue Text"
	magentaText := "this is a magenta Text"
	cyanText := "this is a cyan Text"
	yellowText := "this is a yellow Text"

	fmt.Println(color.Text(redText, color.Red))
	fmt.Println(color.Text(greenText, color.Green))
	fmt.Println(color.Text(blueText, color.Blue))
	fmt.Println(color.Text(magentaText, color.Magenta))
	fmt.Println(color.Text(cyanText, color.Cyan))
	fmt.Println(color.Text(yellowText, color.Yellow))

	// multiple colors
	fmt.Println(color.Text("this is a bold red Text", color.Red, color.Bold))

}
