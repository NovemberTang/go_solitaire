package main

import (
	"fmt"
	. "solitaire/cards"
)

func main() {

	// deck := createShuffledDeck()
	heartStack := &SuitStack{Suit: Hearts, StackHead: nil}
	heartStack.Push(Card{Value: Ace, Suit: Hearts})
	_, err := heartStack.Push(Card{Value: Ace, Suit: Hearts})
	if err != nil {
		fmt.Println(err)
	}

}
