package cards

import "math/rand"

type Colour string

const (
	Red   Colour = "red"
	Black Colour = "black"
)

type Suit struct {
	Name   string
	Colour Colour
}

var (
	Hearts   = Suit{Name: "hearts", Colour: Red}
	Spades   = Suit{Name: "spades", Colour: Black}
	Diamonds = Suit{Name: "diamonds", Colour: Red}
	Clubs    = Suit{Name: "clubs", Colour: Black}
)

type Value int

const (
	Ace Value = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Value Value
	Suit  Suit
}

func createShuffledDeck() []Card {
	deck := []Card{}
	for i := Ace; i <= King; i++ {
		deck = append(deck, Card{Value: i, Suit: Hearts})
		deck = append(deck, Card{Value: i, Suit: Spades})
		deck = append(deck, Card{Value: i, Suit: Diamonds})
		deck = append(deck, Card{Value: i, Suit: Clubs})
	}

	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	return deck
}
