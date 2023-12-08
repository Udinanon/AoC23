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

// store data aout hands to make them easy to order
type hand struct {
	cards [5]card
	rank  rank
	bid   int
}

// need to define methods for Sort interface
type byRanking []hand

// trivial
func (list byRanking) Len() int {
	return len(list)
}

// trivial
func (list byRanking) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// implements the ordering based on rank and individual card vakues as fallback
func (list byRanking) Less(i, j int) bool {
	if list[i].rank == list[j].rank { // if rank is equal
		for k, char := range list[i].cards { // check all cards
			if char == list[j].cards[k] {
				continue
			}
			return char < list[j].cards[k]
		}
	}
	return list[i].rank < list[j].rank
}

// make sit esy to get card values from string
var card_map = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

// makes it easy to use enums
type rank int

// enums on rank
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
	// thanks to the interface we definded we can just ask Go to sort them for us
	sort.Sort(byRanking(hands))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// and compute the result
	result := 0
	for i, hand := range hands {
		result += hand.bid * (i + 1)
	}
	// return result
	fmt.Println(result)
}

func parseLine(line string) hand {
	// we start by simply parsing the string into values and parts
	split_line := strings.Split(line, " ")
	new_bid, _ := strconv.Atoi(split_line[1])
	new_cards := []card{}
	for _, char := range strings.Split(split_line[0], "") {
		new_card := card(card_map[char])
		new_cards = append(new_cards, new_card)
	}
	// here the magic happens
	new_rank := compute_rank(new_cards)
	new_hand := hand{bid: new_bid, cards: [5]card(new_cards), rank: new_rank}
	return new_hand
}

func compute_rank(cards []card) rank {
	// this keeps track of which card appears how many times
	value_map := map[card]int{}
	for _, card := range cards {
		value_map[card]++
	}
	// we don' care which cards they actually are, only how much they appear
	values := []int{}
	for _, value := range value_map {
		values = append(values, value)
	}
	// and we want to know what are the max values
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	// and we use those to get the rank
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
