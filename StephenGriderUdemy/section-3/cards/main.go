package main

// import "fmt"

func main() {
	cards := newDeck()
	// cards := newDeckFromFile("myCards")
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	cards.shuffle()
	cards.print()
	// greeting := "Hey!"
	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())
	// cards.saveToFile("myCards")

}
