package cards

import (
	"fmt"
)

type Stack struct { //todo make generic
	card     Card
	nextCard *Stack
}

type SuitStack struct {
	Suit      Suit //could be private?
	StackHead *Stack
	size      int
}

func addToStack(s *SuitStack, card Card) (*SuitStack, error) {
	s.StackHead = &Stack{card: card, nextCard: s.StackHead}
	s.size++
	fmt.Println("Card pushed to stack")
	return s, nil
}

func (s *SuitStack) Push(card Card) (*SuitStack, error) {
	//card must be of the same suit, and one value higher than the current top card
	stackIsFull := s.size == 13
	if stackIsFull {
		return s, fmt.Errorf("Stack is full")
	}

	cardMatchesSuit := card.Suit == s.Suit
	if !cardMatchesSuit {
		return s, fmt.Errorf("cannot add a %s card to a %s stack", card.Suit.Name, s.Suit.Name)
	}

	nonAceToEmptyStack := s.StackHead == nil && card.Value != Ace
	if nonAceToEmptyStack {
		return s, fmt.Errorf("cannot place a %d card on an empty stack", card.Value)
	}

	invalidNextCardValue := s.StackHead != nil && card.Value-s.StackHead.card.Value != 1
	if invalidNextCardValue {
		return s, fmt.Errorf("cannot place a %d value card on top of a %d value card", card.Value, s.StackHead.card.Value)
	}

	addToStack(s, card)
	fmt.Println(s.StackHead.card)
	return s, nil
}

type Foundation struct {
	Hearts   SuitStack
	Spades   SuitStack
	Diamonds SuitStack
	Clubs    SuitStack
}

func (s *Foundation) Push(card Card) (*Foundation, error) {
	switch card.Suit {
	case Hearts:
		_, err := s.Hearts.Push(card)
		if err != nil {
			return s, err
		}
	case Spades:
		_, err := s.Spades.Push(card)
		if err != nil {
			return s, err
		}
	case Diamonds:
		_, err := s.Diamonds.Push(card)
		if err != nil {
			return s, err
		}
	case Clubs:
		_, err := s.Clubs.Push(card)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}
