package main

import "fmt"

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("color hex %v is %v\n", color, hex)
	}
}
func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	colors["blue"] = "#0000ff"

	delete(colors, "blue")

	colors["blue"] = "#0000ff"

	fmt.Println(colors)
	printMap(colors)
}
