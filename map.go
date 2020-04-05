package main

import "fmt"

func main() {
	//     colors := map[string]string{
	//         "red":"#ff0000",
	//         "green":"#4bf745",
	//  }
	//var colors map[string]string
	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "#4bf745"
	//delete(colors,"red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
