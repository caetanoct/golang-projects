package main

import "fmt"

func main() {
	deck, err := newDeckFromFile("deck.txt")
	if err == nil {
		deck.print()
		deck.shuffle()
		deck.print()
	} else {
		fmt.Println(err)
	}
}
