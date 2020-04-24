package v1

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOrderBy(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  string
	}{
		{name: "Ordered already", hand: "1H 2H 3H 4H 5H", result: "1H 2H 3H 4H 5H"},
		{name: "Shuffled 1", hand: "1H 7H 3H 6H 2H", result: "1H 2H 3H 6H 7H"},
		{name: "Shuffled face cards", hand: "TS KS QS AS JS", result: "TS JS QS KS AS"},
		{name: "All cards same", hand: "TS TS TS TS TS", result: "TS TS TS TS TS"},
		{name: "Shuffled 4 duplicate", hand: "TS TS 8S TS TS", result: "8S TS TS TS TS"},
		{name: "Shuffled 3 duplicate", hand: "TS TS 8S 4S TS", result: "4S 8S TS TS TS"},
		{name: "Shuffled 2 duplicate", hand: "TS TS 8S 4S 3S", result: "3S 4S 8S TS TS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.orderHand(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == strings.Join(result, " "))
		})
	}
}

func TestIsStraight(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "1-5, hearts", hand: "1H 2H 3H 4H 5H", result: true},
		{name: "1-6, 5 excluded", hand: "1H 2H 3H 4H 6H", result: false},
		{name: "10-A, hearts", hand: "TS JS QS KS AS", result: true},
		{name: "K-4, spades, invalid wrap around", hand: "KS AS 2S 3S 4S", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isStraight(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
		})
	}
}

func TestIsFourOfAKind(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "4 of a kind", hand: "AH AH AH AH 5H", result: true},
		{name: "not 4 of a kind", hand: "AH 2H AH AH 6H", result: false},
		{name: "not 4 of a kind", hand: "TS JS QS KS AS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isFourOfAKind(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
		})
	}
}

func TestIsThreeOfAKind(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "3 of a kind", hand: "AH AH AH 9H 5H", result: true},
		{name: "not 3 of a kind", hand: "2H 6H 6H AH AH", result: false},
		{name: "not 3 of a kind", hand: "TS JS QS KS AS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isThreeOfAKind(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
		})
	}
}

func TestIsTwoPair(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "2 pair", hand: "AH AH 5H 5H 6H", result: true},
		{name: "2 pair", hand: "AH AH 6H 5H 5H", result: true},
		{name: "2 pair", hand: "5S JS JS KS KS", result: true},
		{name: "not 2 pair", hand: "5S JS JS QS KS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isTwoPair(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
		})
	}
}

func TestIsPair(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		i		int
		hand    string
		result  bool
	}{
		{name: "Pair", i: 14, hand: "AH AH 5H 7H 6H", result: true},
		{name: "Pair", i: 14, hand: "KH AH AH 5H 2H", result: true},
		{name: "Pair", i: 11, hand: "5S QS JS JS KS", result: true},
		{name: "Pair", i: 11, hand: "5S QS KS JS JS", result: true},
		{name: "not 2 pair", i: 0, hand: "5S 8S JS QS KS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, i := p.isPair(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
			assert.True(t, i == tt.i)
		})
	}
}

func TestIsFlush(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "Flush", hand: "AH AH 5H 7H 6H", result: true},
		{name: "Not flush", hand: "KH AS AH 5H 2H", result: false},
		{name: "Not flush", hand: "5S QD JS JC KS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isFlush(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
		})
	}
}

func TestIsFullHouse(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "Full House", hand: "AH AH AH 7H 7H", result: true},
		{name: "Full House", hand: "AH AH 7H 7H 7H", result: true},
		{name: "Full house", hand: "KH AS AH AC KH", result: true},
		{name: "Not full house", hand: "5S QD JS JC KS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isFullHouse(p.orderHand(strings.Split(tt.hand, " ")))
			assert.True(t, tt.result == result)
		})
	}
}

func TestIsHighCard(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand    string
		result  bool
	}{
		{name: "High Card", hand: "AH AH AH 7H 7H", result: true},
		{name: "High Card", hand: "AH AH 7H 7H 7H", result: true},
		{name: "High Card", hand: "KH AS AH AC KH", result: true},
		{name: "High Card", hand: "5S QD JS JC KS", result: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.isHighCard(strings.Split(tt.hand, " "))
			assert.True(t, tt.result == result)
		})
	}
}

func TestCompareHands(t *testing.T) {
	p := Poker{}

	tests := []struct {
		name    string
		hand1   string
		hand2	string
		result  bool
	}{
		{name: "4 vs 3", hand1: "AH AH AH AH 7H", hand2: "AH AH AH KH QH", result: true},
		{name: "4 vs 2", hand1: "AH AH AH AH 7H", hand2: "AH AH 7H 8H 9H", result: true},
		{name: "Straight Flush vs 4", hand1: "2H 3H 4H 5H 6H", hand2: "KH KS KH KC AH", result: true},
		{name: "Flush vs Straight", hand1: "5S 7S JS QS KS", hand2: "5S 6D 7S 8C 9S", result: true},
		{name: "High Card vs Pair", hand1: "5S QD JS 4C KS", hand2: "5S QD JS JC KS", result: false},
		{name: "Pair vs Pair", hand1: "5S QD JS JC KS", hand2: "5S QD JS KC KS", result: false},
		{name: "Full House vs Flush", hand1: "5S 5D 5S 6C 6S", hand2: "5S QS JS JC KS", result: true},
		{name: "Flush vs Flush", hand1: "9S 8S 7S 6S 5S", hand2: "9D 8D 7D 6D 4D", result: true},
		{name: "Trip vs Trip", hand1: "5S 5D 5S JC KS", hand2: "5S KD JS JC JS", result: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.compareHands(tt.hand1, tt.hand2)
			assert.True(t, tt.result == result)
		})
	}
}
