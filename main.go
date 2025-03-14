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

type Stack struct { //todo make generic
	card     Card
	nextCard *Stack
}

type SuitStack struct {
	suit      Suit
	stackHead *Stack
}

func (s *SuitStack) push(card Card) (*SuitStack, error) {
	//card must be of the same suit, and one value higher than the current top card
	if card.Suit != s.suit {
		return s, fmt.Errorf("Card suit does not match stack suit")
	}

	if s.stackHead == nil {
		if card.Value != Ace {
			return s, fmt.Errorf("first card in stack must be an Ace")
		} else {
			s.stackHead = &Stack{card: card, nextCard: nil}
			return s, nil
		}
	}

	if card.Value != s.stackHead.card.Value+1 {
		return s, fmt.Errorf("current card value is %d. The next card must be %d, not %d", s.stackHead.card.Value, s.stackHead.card.Value+1, card.Value)
	}

	s.stackHead = &Stack{card: card, nextCard: s.stackHead}
	fmt.Println("Card pushed to stack")
	return s, nil
}

func main() {

	// deck := createShuffledDeck()

	heartStack := &SuitStack{suit: Hearts, stackHead: nil}
	heartStack.push(Card{Value: Ace, Suit: Hearts})
	_, err := heartStack.push(Card{Value: Ace, Suit: Hearts})
	if err != nil {
		fmt.Println(err)
	}

}
