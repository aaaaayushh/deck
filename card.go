//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // technically not a suit, but it's a different type of card which some games have uses for, hence adding here
)

type Rank uint8

const (
	Ace Rank = iota + 1
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
	// any card will have 1. a suit (hearts,diamonds,clubs,spades) and 2. a rank (ace, 2, 3, ..., J, Q, K)
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
