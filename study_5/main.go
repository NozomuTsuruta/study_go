package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot)  {
// 	fmt.Println(eb.getFreeting())
// }

// func printGreeting(sb spanishBot)  {
// 	fmt.Println(sb.getFreeting())
// }

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
