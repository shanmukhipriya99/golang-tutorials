package main

import "fmt"

func main() {
	// var colors map[string]string
	
	// colors := make(map[string]string) //in-built make function

	colors := map[string]string{
		"red": "1",
		"green": "2",//add , after every single prop even the last one
	}

	colors["white"] = "0" //maps use [] not .
	// delete(colors, "red") //in-built delete function
	
	// fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, value := range c {
		fmt.Println(color+"_"+value)
	}
}