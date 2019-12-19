package main

import "fmt"

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
func main() {

	/*
		=============
		=  GO MAPS  =
		=============
	*/

	// METHOD 1
	//var colors map[string]string

	// METHOD 2
	//colors := make(map[string]string)

	// METHOD 3
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

	// ASSIGNING VALUES TO KEYS
	//colors["white"] = "#ffffff"

	// DELETING A KEY
	//delete(colors, "white")
	//fmt.Println(colors)
}
