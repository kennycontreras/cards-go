package main

import "fmt"

func main() {

	cards := newDeck()

	cards.print()
	fmt.Println("here ")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
