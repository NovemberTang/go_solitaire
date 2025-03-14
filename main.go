package main

import (
	"fmt"
	"math/rand"
)

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
	Ace Value = iota
	One
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

type card struct {
	Value Value
	Suit  Suit
}

func createShuffledDeck() []card {
	deck := []card{}
	for i := Ace; i <= King; i++ {
		deck = append(deck, card{Value: i, Suit: Hearts})
		deck = append(deck, card{Value: i, Suit: Spades})
		deck = append(deck, card{Value: i, Suit: Diamonds})
		deck = append(deck, card{Value: i, Suit: Clubs})
	}

	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	return deck
}

func main() {

	deck := createShuffledDeck()

	fmt.Println(deck[:3])
}
