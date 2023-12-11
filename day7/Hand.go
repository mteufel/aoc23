package day7

import (
	"aoc23/util"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	WithoutJoker bool = false
	WithJoker    bool = true
)

type Hand struct {
	Cards     string
	Bid       int
	WithJoker bool
}

type Card struct {
	Label    string
	Strength int
}

func (hand *Hand) GetHighestStrength() (string, int) {
	m := hand.GetStrengths()
	return util.RankMapStringInt(m)[0], m[util.RankMapStringInt(m)[0]]
}

func (hand *Hand) GetStrengths() map[string]int {
	m := make(map[string]int)
	for _, pattern := range hand.CardsAll() {
		if m[pattern.Label] >= 1 {
			m[pattern.Label] = m[pattern.Label] + pattern.Strength + 100000
		} else {
			m[pattern.Label] = pattern.Strength
		}
	}
	n := make(map[string]int)
	for _, index := range util.RankMapStringInt(m) {
		n[index] = m[index]
	}
	return m
}

func (hand *Hand) ContainsJoker() bool {
	return strings.Contains(hand.Cards, "J")
}

func (hand *Hand) StrongerThan(card string) bool {
	compareHand := NewHand(card, -1, hand.WithJoker)
	if hand.GetType() != compareHand.GetType() {
		return false
	}
	r1 := make([]bool, 0)
	for idx, card := range hand.CardsAll() {
		that := compareHand.CardsAll()[idx].Strength
		this := card.Strength
		if this > that {
			r1 = append(r1, true)
		} else {
			r1 = append(r1, false)
		}
	}
	r2 := make([]bool, 0)
	for idx, card := range compareHand.CardsAll() {
		that := hand.CardsAll()[idx].Strength
		this := card.Strength
		if this > that {
			r2 = append(r2, true)
		} else {
			r2 = append(r2, false)
		}
	}

	stronger := ""
	for i, b := range r1 {
		if b != (r2[i]) {
			if b == true {
				stronger = hand.Cards
				break
			} else {
				stronger = compareHand.Cards
				break
			}
		}
	}
	return hand.Cards != stronger
}

func (hand *Hand) GetType() Type {
	result := GetTypeByName("High card")

	cards := hand.Cards
	if hand.WithJoker == WithJoker && hand.ContainsJoker() {
		strongestPattern, _ := hand.GetHighestStrength()
		hand.Cards = strings.ReplaceAll(hand.Cards, "J", strongestPattern)
	}

	cardsDistinct := hand.CardsDistinct()
	for _, pattern := range cardsDistinct {
		index := Search(hand.Cards, pattern.Label)
		count := len(index)
		switch count {
		case 5:
			result = GetTypeByName(FiveOfAKind)
		case 4:
			result = GetTypeByName(FourOfAKind)
		case 3:
			switch result.Name {
			case OnePair:
				result = GetTypeByName(FullHouse)
			default:
				result = GetTypeByName(ThreeOfAKind)
			}
		case 2:
			switch result.Name {
			case ThreeOfAKind:
				result = GetTypeByName(FullHouse)
			case OnePair:
				result = GetTypeByName(TwoPair)
			default:
				result = GetTypeByName(OnePair)
			}
		}
	}
	hand.Cards = cards
	return result
}

/*
func (hand *Hand) CardsDistinct(cards string) []Card {
	_cardsAll := MakeCardsArray(hand.Cards, hand.WithJoker)
	_cardsDistinct := make([]Card, 0)
	for _, c := range _cardsAll {
		if !slices.Contains(_cardsDistinct, c) {
			_cardsDistinct = append(_cardsDistinct, c)
		}
	}
	sort.Slice(_cardsDistinct, func(i, j int) bool {
		return _cardsDistinct[i].Strength < _cardsDistinct[j].Strength
	})
	return _cardsDistinct
}
*/

func (hand *Hand) CardsDistinct() []Card {
	_cardsAll := hand.CardsAll()
	_cardsDistinct := make([]Card, 0)
	for _, c := range _cardsAll {
		if !slices.Contains(_cardsDistinct, c) {
			_cardsDistinct = append(_cardsDistinct, c)
		}
	}
	sort.Slice(_cardsDistinct, func(i, j int) bool {
		return _cardsDistinct[i].Strength < _cardsDistinct[j].Strength
	})
	return _cardsDistinct
}

func (hand *Hand) CardsAll() []Card {
	return MakeCardsArray(hand.Cards, hand.WithJoker)
}

func MakeCardsArray(cards string, withJoker bool) []Card {
	_cards := make([]Card, 0)
	for _, s := range strings.Split(cards, "") {
		strength := 0
		switch s {
		case "A":
			strength = 14
		case "K":
			strength = 13
		case "Q":
			strength = 12
		case "J":
			if withJoker {
				strength = 0
			} else {
				strength = 11
			}
		case "T":
			strength = 10
		case "9":
			strength = 9
		case "8":
			strength = 8
		case "7":
			strength = 7
		case "6":
			strength = 6
		case "5":
			strength = 5
		case "4":
			strength = 4
		case "3":
			strength = 3
		case "2":
			strength = 2

		}
		_cards = append(_cards, Card{Label: s, Strength: strength})
	}
	return _cards
}

func NewHand(cards string, bid int, withJoker bool) Hand {
	return Hand{Cards: cards, Bid: bid, WithJoker: withJoker}
}

func Search(text string, pattern string) []int {

	// implementation of "Naive Search" pattern
	// see (german) https://de.wikipedia.org/wiki/String-Matching-Algorithmus

	results := make([]int, 0)
	lengthText := len(text)
	lengthPattern := len(pattern)

	for i := range N(lengthText - lengthPattern + 1) {
		j := 0
		for j < lengthPattern {
			if strings.Split(text, "")[i+j] != strings.Split(pattern, "")[j] {
				break

			}
			j++
			if j == lengthPattern {
				results = append(results, i)
			}
		}
	}
	return results
}

func N(size int) []struct{} {
	return make([]struct{}, size)
}

func Load(file string, withJoker bool) []Hand {
	lines := util.ReadFile("/day7/" + file)

	hands := make([]Hand, 0)

	// -------------------------------
	// Create Hand
	// ------------------------------------
	for _, line := range lines {
		cards := strings.Split(line, " ")[0]
		bid, _ := strconv.Atoi(strings.Split(line, " ")[1])
		hands = append(hands, NewHand(cards, bid, withJoker))

	}

	// -------------------------------
	// Sort by Type and Strongness
	// ------------------------------------
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.GetType().Strength >= b.GetType().Strength {
			if a.StrongerThan(b.Cards) {
				return -1
			} else {
				return 0
			}

		}
		return -1
	})

	return hands
}
