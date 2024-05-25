package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	/*
		The selected code is using the fmt.Println function to print the Card object.
		In Go, when the fmt package's print functions (like fmt.Println, fmt.Printf, etc.) encounter a value of a type that
		satisfies the fmt.Stringer interface, they automatically call the String method on that value to get its string
		representation.
	*/
	// Test Ace of Hearts
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	// Test Two of Spades
	fmt.Println(Card{Rank: Two, Suit: Spade})

	// Test King of Diamonds
	fmt.Println(Card{Rank: King, Suit: Diamond})

	// Test Queen of Clubs
	fmt.Println(Card{Rank: Queen, Suit: Club})

	// Test Joker
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Spades
	// King of Diamonds
	// Queen of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Errorf("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	// First card should be Ace of Clubs
	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
		t.Errorf("Expected Ace of Spades, received %v", cards[0])
	}

	// Last card should be King of Spades
	if cards[len(cards)-1] != (Card{Rank: King, Suit: Heart}) {
		t.Errorf("Expected King of Hearts, received %v", cards[len(cards)-1])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	// First card should be Ace of Clubs
	if cards[0] != (Card{Rank: Ace, Suit: Spade}) {
		t.Errorf("Expected Ace of Spades, received %v", cards[0])
	}

	// Last card should be King of Spades
	if cards[len(cards)-1] != (Card{Rank: King, Suit: Heart}) {
		t.Errorf("Expected King of Hearts, received %v", cards[len(cards)-1])
	}
}
func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Expected 3 Jokers, received %d", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Errorf("Expected to filter out Twos and Threes")
		}
	}
}
