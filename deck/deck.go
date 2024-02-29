package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(d)
	for i := 0; i < length; i++ {
		index := r.Int() % length
		d[i], d[index] = d[index], d[i]
	}
}

func newDeckFromFile(filename string) (deck, error) {
	bytearray, err := os.ReadFile(filename)
	if err == nil {
		return deck(strings.Split(string(bytearray), ",")), nil
	} else {
		return nil, err
	}
}
func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func deal(d deck, amount int) (deck, deck) {
	return d[:amount], d[amount:]
}

// d is the actual copy of the deck
// deck is the type which can call this method
func (d deck) print() {
	for i, card := range d {
		fmt.Println("index=", i, "card=", card)
	}
}
