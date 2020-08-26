package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// := using only when declare an variable
	card := newCard()

	fmt.Println(card)

}

func newCard() string {
	return "Five of Diamonds"
}
