package cards

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	deck := createShuffledDeck()
	//assert that there are 52 cards in the deck
	if len(deck) != 52 {
		t.Errorf("Expected deck to have 52 cards, but got %d", len(deck))
	}
	//assert that there are 13 cards of each suit
	hearts := 0
	spades := 0
	diamonds := 0
	clubs := 0
	for _, card := range deck {
		switch card.Suit {
		case Hearts:
			hearts++
		case Spades:
			spades++
		case Diamonds:
			diamonds++
		case Clubs:
			clubs++
		}
	}
	if hearts != 13 {
		t.Errorf("Expected 13 hearts, but got %d", hearts)
	}
	if spades != 13 {
		t.Errorf("Expected 13 spades, but got %d", spades)
	}
	if diamonds != 13 {
		t.Errorf("Expected 13 diamonds, but got %d", diamonds)
	}
	if clubs != 13 {
		t.Errorf("Expected 13 clubs, but got %d", clubs)
	}
}
