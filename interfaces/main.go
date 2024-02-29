package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type portugueseBot struct{}

func main() {
	eb := englishBot{} // english bot satisfies bot interface, so it is accepted into printGreeting
	pb := portugueseBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (portugueseBot) getGreeting() string {
	return "Oi!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(b englishBot) {
// 	fmt.Println(b.getGreeting())
// }

// func printGreeting(b portugueseBot) {
// 	fmt.Println(b.getGreeting())
// }
