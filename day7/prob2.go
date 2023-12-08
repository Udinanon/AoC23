package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card int

type hand struct {
	cards [5]card
	rank  rank
	bid   int
}

type byRanking []hand

func (list byRanking) Len() int {
	return len(list)
}

func (list byRanking) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list byRanking) Less(i, j int) bool {
	if list[i].rank == list[j].rank {
		for k, char := range list[i].cards {
			if char == list[j].cards[k] {
				continue
			}
			return char < list[j].cards[k]
		}
	}
	return list[i].rank < list[j].rank
}

var card_map = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	//	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1, // we moved the joker to it's value, use for fallback ordering
}

type rank int

const (
	HighCard rank = iota + 1 //1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	// open file
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	hands := []hand{}

	for scanner.Scan() {

		new_hand := parseLine(scanner.Text())
		hands = append(hands, new_hand)
	}

	sort.Sort(byRanking(hands))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	result := 0
	for i, hand := range hands {
		result += hand.bid * (i + 1)
	}
	// return result
	fmt.Println(result)
}

func parseLine(line string) hand {
	split_line := strings.Split(line, " ")
	new_bid, _ := strconv.Atoi(split_line[1])
	new_cards := []card{}
	for _, char := range strings.Split(split_line[0], "") {
		new_card := card(card_map[char])
		new_cards = append(new_cards, new_card)
	}
	new_rank := compute_rank(new_cards)
	new_hand := hand{bid: new_bid, cards: [5]card(new_cards), rank: new_rank}
	return new_hand
}

func compute_rank(cards []card) rank {
	value_map := map[card]int{}
	jokers := 0
	for _, card := range cards {
		if card == 1 { // Joker flag
			jokers++ // we need to count the jokers for later
			continue
		}
		value_map[card]++
	}
	values := []int{}
	for _, value := range value_map {
		values = append(values, value)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	if len(values) == 0 { // happens if a hand is ONLY jokers
		values = []int{0}
	}
	values[0] += jokers // it's notied that jokers always have the most impact if they add to the most common card, TwoPairs and FullHouses are never as good as Three or Fouor
	switch values[0] {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if values[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if values[1] == 2 {
			return TwoPair
		}
		return OnePair
	default:
		return HighCard
	}

}
