package deck

import "fmt"

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
