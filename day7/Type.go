package day7

import "slices"

const (
	FiveOfAKind  string = "Five of a kind"
	FourOfAKind  string = "Four of a kind"
	FullHouse    string = "Full house"
	ThreeOfAKind string = "Three of a kind"
	TwoPair      string = "Two Pair"
	OnePair      string = "One Pair"
	HighCard     string = "High card"
)

type Type struct {
	Name     string
	Strength int
}

var types = []Type{
	{Name: FiveOfAKind, Strength: 70000},
	{Name: FourOfAKind, Strength: 60000},
	{Name: FullHouse, Strength: 50000},
	{Name: ThreeOfAKind, Strength: 40000},
	{Name: TwoPair, Strength: 30000},
	{Name: OnePair, Strength: 20000},
	{Name: HighCard, Strength: 10000},
}

func GetTypeByName(name string) Type {
	idx := slices.IndexFunc(types, func(t Type) bool {
		return t.Name == name
	})
	return types[idx]
}
