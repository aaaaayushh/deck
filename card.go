//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // SPECIAL CASE technically not a suit, but it's a different type of card which some games have uses for, hence adding here
)

/*
[...]Suit is an array type. The ... is a special syntax in Go that automatically determines the length of the array
based on the number of elements provided in the array literal.
*/
var suits = [...]Suit{Spade, Diamond, Club, Heart}

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

const (
	minRank = Ace
	maxRank = King
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

// New functional options used here
// this is done to keep a function that is supposed to be used by a client extensible and flexible
// refer: https://golang.cafe/blog/golang-functional-options-pattern.html
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort this function can be passed as a parameter to the New function
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort Custom sort function for user
// func(cards []Card) func(i, j int) bool -> this is the type of the less function (look at Less function definition)
// func([]Card) []Card -> this is the type of the Sort function itself
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less refer to https://cs.opensource.google/go/go/+/go1.22.3:src/sort/slice.go;l=24
// returns a function (i,j) which in turn returns a boolean value
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// to sort the deck in order of Spade(1,2...K),Diamond(...) and so on
func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	// creating a new random number generator that is seeded with the current time.
	// The variable r can then be used to generate random numbers
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	// perm will be of the form [0,4,2,3,1...]
	// "i" is the index of the array i.e. 0,1,2,3,...
	// "j" is the value at that index of the perm array
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers add n number of jokers to the deck
// called like New(Jokers(3)) -- adds 3 jokers to the deck
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker, Rank: Rank(i)})
		}
		return cards
	}
}

// Filter function to filter out certain cards from the deck
// f is a function that returns true or false depending on whether a card should be in the deck or not
// f is provided by the user according to their needs
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// Deck function to generate more decks
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
