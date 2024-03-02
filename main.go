package main

import (
	"fmt"
    "github.com/Jett65/52Cards/cards"
)

func main() {
	deck := cards.NewBlackJackDeck()
    for i := 0; i < 50000; i++ {
        deck = deck.Shuffle()
    }

    fmt.Println(deck)

}
