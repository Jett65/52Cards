package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Card struct {
    value         string
	sute          string
	numaric_value int
}

func (card *Card) IsBlank() bool {
    if card.sute == "" && card.value == "" && card.numaric_value == 0 {
        return true
    }

    return false
}

type Deck struct {
	cards []Card
}

func (deck *Deck) IsEmpty() bool {
	return len(deck.cards) == 0
}

func (deck *Deck) Add(card Card) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) Shuffle() Deck {
    var shuffled [52]Card
    cards := deck.cards

    for _, v := range cards { 
        cardex := rand.Intn(len(cards))
        for {
            if shuffled[cardex].IsBlank() {
                break
            }

            cardex = rand.Intn(len(cards))
        }
        
        shuffled[cardex] = v
    }

    return Deck {
        cards: shuffled[0:],
    }
}

func (deck *Deck) Deal() Card {
	if deck.IsEmpty() {
		return Card{}
	}

	card := deck.cards[len(deck.cards)-1]
	deck.cards = deck.cards[:len(deck.cards)-1]

	return card
}

func (deck *Deck) Peek() Card {
	if deck.IsEmpty() {
		return Card{}
	}

	return deck.cards[len(deck.cards)-1]
}

// Create Decks that are filled
func New52Deck() Deck {
	var deck Deck
	sutes := []string{"spades", "harts", "clubs", "diamonds"}
	values := []string{"Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	for _, sute := range sutes {
		for _, value := range values {
			deck.Add(Card{
				value: value,
                sute:  sute,
			})
		}
	}

	return deck
}

func NewBlackJackDeck() Deck {
	var deck Deck
	sutes := []string{"spades", "harts", "clubs", "diamonds"}
	values := []string{"Ace", "King", "Queen", "Jack", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	for _, sute := range sutes {
		for _, value := range values {
            var num_val int

            if value == "Ace" {
                num_val = 11
            } else if value == "King" || value == "Queen" || value == "Jack" {
                num_val = 10
            } else {
                num_val, _ = strconv.Atoi(value)
            }

			deck.Add(Card{
                value: value,
				sute:  sute,
                numaric_value: num_val,
			})
		}
	}

    return deck
}

func main() {
	cards := NewBlackJackDeck() 
    for i := 0; i < 50000; i++ {
        cards = cards.Shuffle()
    }

    fmt.Println(cards)

}
