package main

import "fmt"

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

func main() {

	myCard := card{
		Value: Three,
		Suit:  Spades,
	}

	fmt.Println(myCard)
}
