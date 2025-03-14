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
}

func (s *SuitStack) Push(card Card) (*SuitStack, error) {
	//card must be of the same suit, and one value higher than the current top card
	if card.Suit != s.Suit {
		return s, fmt.Errorf("Card suit does not match stack suit")
	}

	if s.StackHead == nil {
		if card.Value != Ace {
			return s, fmt.Errorf("first card in stack must be an Ace")
		} else {
			s.StackHead = &Stack{card: card, nextCard: nil}
			return s, nil
		}
	}

	if card.Value != s.StackHead.card.Value+1 {
		return s, fmt.Errorf("current card value is %d. The next card must be %d, not %d", s.StackHead.card.Value, s.StackHead.card.Value+1, card.Value)
	}

	s.StackHead = &Stack{card: card, nextCard: s.StackHead}
	fmt.Println("Card pushed to stack")
	return s, nil
}
