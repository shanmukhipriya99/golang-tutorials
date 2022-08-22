package main

import "fmt"

type bot interface {
	getGreeting() string
}

type eBot struct {}
type sBot struct {}

func main() {
	eb := eBot{}
	sb := sBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb eBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb sBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eBot) getGreeting() string { //not using eb eBot as we're not gonna use eb anywhere
	return "Hi there!"
}

func (sBot) getGreeting() string {
	return "Hola!"
}