package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colours map[string]string		// another way to create a map
	colours := make(map[string]string) // yet another way

	fmt.Println(colors)
	fmt.Println(colours)

	colours["white"] = "#ffffff"
	fmt.Println(colours)

	delete(colours, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
