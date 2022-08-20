package main

// import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// cards.print()
	// greeting := "Hey!"
	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())
	cards.saveToFile("myCards")
}
