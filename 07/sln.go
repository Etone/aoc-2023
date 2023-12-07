package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandValue = int

const (
	FiveOfAKind  HandValue = 6
	FourOfAKind  HandValue = 5
	FullHouse    HandValue = 4
	ThreeOfAKind HandValue = 3
	TwoPair      HandValue = 2
	Pair         HandValue = 1
	Highest      HandValue = 0
)

type Player struct {
	cards Hand
	bid   int
}

type Hand struct {
	cards []Card
	value HandValue
}

type Card = rune

var cardOrder = map[Card]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

//go:embed input.txt
var input string

func parsePlayer1(line string) (player Player) {
	cards, bid, _ := strings.Cut(line, " ")
	player.bid, _ = strconv.Atoi(bid)
	player.cards.cards = []rune(cards)
	player.cards.value = calculateValue(player.cards.cards)

	return
}

func parsePlayer2(line string) (player Player) {
	cards, bid, _ := strings.Cut(line, " ")
	player.bid, _ = strconv.Atoi(bid)
	player.cards.cards = []rune(cards)
	player.cards.value = calculateValue2(player.cards.cards)

	return
}

func calculateValue(hand []rune) HandValue {
	repeated := make([]int, 5)
	cards := string(hand)
	for _, card := range hand {
		repeatsOfCard := strings.Count(cards, string(card))
		repeated[repeatsOfCard-1] += 1
	}

	if repeated[4] == 5 {
		return FiveOfAKind
	}
	if repeated[3] == 4 {
		return FourOfAKind
	}
	if repeated[2] == 3 {
		if repeated[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if repeated[1] == 4 {
		return TwoPair
	}
	if repeated[1] == 2 {
		return Pair
	}
	return Highest
}

func calculateValue2(hand []rune) HandValue {
	cards := string(hand)
	jokers := strings.Count(cards, "J")
	valueWithoutJoker := calculateValue(hand)

	return valueWithJoker(valueWithoutJoker, jokers)
}

func valueWithJoker(valueWithoutJoker, jokers int) HandValue {
	if jokers == 4 {
		return FiveOfAKind
	}

	if jokers == 3 {
		if valueWithoutJoker == ThreeOfAKind {
			return FourOfAKind
		}
		if valueWithoutJoker == FullHouse {
			return FiveOfAKind
		}
	}

	if jokers == 2 {
		if valueWithoutJoker == Pair {
			// Pair must be the jokers, any other card gives three of a kind
			return ThreeOfAKind
		}
		if valueWithoutJoker == TwoPair {
			// One pair is the jokers, there is an extra pair which means we can get four of a kind
			return FourOfAKind
		}
		if valueWithoutJoker == FullHouse {
			// Three of a Kind + 2 jokers is the only way here
			return FiveOfAKind
		}
	}

	if jokers == 1 {
		if valueWithoutJoker == Highest {
			return Pair
		}
		if valueWithoutJoker == Pair {
			return ThreeOfAKind
		}
		if valueWithoutJoker == TwoPair {
			return FullHouse
		}
		if valueWithoutJoker == ThreeOfAKind {
			return FourOfAKind
		}
		if valueWithoutJoker == FourOfAKind {
			return FiveOfAKind
		}
	}
	return valueWithoutJoker
}

func compareHands(hand1, hand2 Hand) bool {
	if hand1.value == hand2.value {
		for i := range hand1.cards {
			if cardOrder[hand1.cards[i]] == cardOrder[hand2.cards[i]] {
				continue
			}
			return cardOrder[hand1.cards[i]] < cardOrder[hand2.cards[i]]
		}
	}
	return hand1.value < hand2.value
}

func part1() {
	playedGames := []Player{}
	gamesInput := strings.Split(input, "\n")
	for _, handText := range gamesInput {
		playedGames = append(playedGames, parsePlayer1(handText))
	}

	sort.Slice(playedGames, func(i, j int) bool {
		return compareHands(playedGames[i].cards, playedGames[j].cards)
	})

	totalWinning := 0

	for rank, game := range playedGames {
		totalWinning += game.bid * (rank + 1)
	}

	fmt.Printf("Part 1: %d\n", totalWinning)
}

func part2() {
	//change Joker to be least valuable
	cardOrder['J'] = -1

	playedGames := []Player{}
	gamesInput := strings.Split(input, "\n")
	for _, handText := range gamesInput {
		playedGames = append(playedGames, parsePlayer2(handText))
	}

	sort.Slice(playedGames, func(i, j int) bool {
		return compareHands(playedGames[i].cards, playedGames[j].cards)
	})

	totalWinning := 0

	for rank, game := range playedGames {
		totalWinning += game.bid * (rank + 1)
	}

	fmt.Printf("Part 2: %d\n", totalWinning)
}

func main() {
	part1()
	part2()
}
