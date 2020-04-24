package v1

import (
	"log"
	"strconv"
	"strings"
)

type Poker struct {
	s Stats
}

// return true if p1 > p2. Else return false
func (p *Poker) compareHands(p1 string, p2 string) bool{
	p1Cards := p.orderHand(strings.Split(p1, " "))
	p2Cards := p.orderHand(strings.Split(p2, " "))

	p1Val := p.evaluate(p1Cards)
	p2Val := p.evaluate(p2Cards)

	if p1Val > p2Val {
		p.s.Player1Win()
	} else {
		p.s.Player2Win()
	}

	p.s.RoundInc()

	return p1Val > p2Val
}

func (p *Poker) evaluate(hand []string) int {
	if p.isStraightFlush(hand) {
		p.s.StraightFlush()
		return 100000000 * p.getFaceValue(hand[4]) + p.getCardTotalValue(hand[0])
	} else if p.isFourOfAKind(hand) {
		p.s.FourOfAKind()
		if hand[0][0] == hand[1][0] {
			return 10000000 * p.getFaceValue(hand[0]) + p.getCardTotalValue(hand[4])
		}
		return 10000000 * p.getFaceValue(hand[4]) + p.getCardTotalValue(hand[0])
	} else if p.isFullHouse(hand) {
		p.s.FullHouse()
		if hand[0][0] == hand[2][0] {
			return 1000000 * p.getFaceValue(hand[0]) + p.getCardTotalValue(hand[4])
		}
		return 1000000 * p.getFaceValue(hand[4]) + p.getCardTotalValue(hand[0])
	} else if p.isFlush(hand) {
		p.s.Flush()
		return 1000000 * p.getFaceValue(hand[0])
	} else if p.isStraight(hand) {
		p.s.Straight()
		return 10000 * p.getCardTotalValue(hand[0])
	} else if p.isThreeOfAKind(hand) {
		p.s.ThreeOfAKind()
		return 1000 * p.getFaceValue(hand[2])
	} else if p.isTwoPair(hand) {
		p.s.TwoPair()
		return 100
	} else if b, i := p.isPair(hand); b {
		p.s.Pair()
		return 10 * i
	} else {
		// High card
		p.s.HighCard()
		return p.getCardTotalValue(hand[0])
	}
}

func (p *Poker) isStraightFlush(hand []string) bool {
	isStraight := p.isStraight(hand)
	isFlush := p.isFlush(hand)
	return isStraight && isFlush
}

func (p *Poker) isFourOfAKind(hand []string) bool {
	return (hand [0][0] == hand[1][0] && hand[1][0] == hand[2][0] && hand[2][0] == hand[3][0]) ||
		(hand[1][0] == hand[2][0] && hand[2][0] == hand[3][0] && hand [3][0] == hand[4][0])
}

func (p *Poker) isFullHouse(hand []string) bool {
	return (hand[0][0] == hand[2][0] && hand[3][0] == hand[4][0]) ||
		(hand[0][0] == hand[1][0] && hand[2][0] == hand[4][0])
}

func (p *Poker) isFlush(hand []string) bool {
	return hand[0][1] == hand[1][1] &&
		hand[0][1] == hand[2][1] &&
		hand[0][1] == hand[3][1] &&
		hand[0][1] == hand[4][1]
}
func (p *Poker) isStraight(hand []string) bool {
	for i := 1; i < 5; i++ {
		currCard := p.getFaceValue(string(hand[i][0]))
		prevCard := p.getFaceValue(string(hand[i-1][0]))

		if currCard-prevCard!=1 {
			return false
		}
	}
	return true
}

func (p *Poker) isThreeOfAKind(hand []string) bool {
	return hand [0][0] == hand[2][0] || hand[1][0] == hand[3][0] || hand[2][0] == hand[4][0]

}

func (p *Poker) isTwoPair(hand []string) bool {
	return (hand [0][0] == hand[1][0] && hand[2][0] == hand[3][0]) ||
	(hand [0][0] == hand[1][0] && hand[3][0] == hand[4][0]) ||
	(hand [1][0] == hand[2][0] && hand[3][0] == hand[4][0])
}

func (p *Poker) isPair(hand []string) (bool, int) {
	if hand [0][0] == hand[1][0] {
		return true, p.getFaceValue(hand[0])
	} else if hand [1][0] == hand[2][0] {
		return true, p.getFaceValue(hand[1])
	} else if hand [2][0] == hand[3][0] {
		return true, p.getFaceValue(hand[2])
	} else if hand [3][0] == hand[4][0] {
		return true, p.getFaceValue(hand[3])
	} else {
		return false, 0
	}
}

func (p *Poker) isHighCard(hand []string) bool {
	return true
}

func (p *Poker) getFaceValue(card string) int {
	val, err := strconv.Atoi(string(card[0]))
	if err!=nil {
		switch string(card[0]) {
		case "T":
			val = 10
		case "J":
			val = 11
		case "Q":
			val = 12
		case "K":
			val = 13
		case "A":
			val = 14
		default:
			log.Fatalf("Unable to parse card: %v", card)
		}
	}

	return val
}

func (p *Poker) getSuiteValue(card string) int {
	var val int
	switch string(card[1]) {
	case "S":
		val = 4
	case "H":
		val = 3
	case "D":
		val = 2
	case "C":
		val = 1
	}
	return val
}

func (p *Poker) getCardTotalValue(card string) int {
	return p.getSuiteValue(card) + p.getFaceValue(card)
}



func (p *Poker) orderHand(hand []string) []string {
	t := make([]string, 5, 5)
	for i := 0; i < 5; i++ {
		curr := p.getFaceValue(hand[i])
		pos := 0
		for j := 0; j < 5; j++ {
			if i==j {
				continue
			}

			checkCard := p.getFaceValue(hand[j])

			if j < i {
				if curr > checkCard-1{
					pos++
				}
			} else {
				if curr > checkCard{
					pos++
				}
			}
		}
		t[pos] = hand[i]
	}
	return t
}