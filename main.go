package main

import (
	"fmt"
	. "solitaire/cards"
)

func main() {

	// deck := createShuffledDeck()
	heartStack := &SuitStack{Suit: Hearts, StackHead: nil}
	_, err := heartStack.Push(Card{Value: Ace, Suit: Hearts})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Card pushed to stack")
	}

	_, err = heartStack.Push(Card{Value: Ace, Suit: Hearts})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Card pushed to stack")
	}

}
