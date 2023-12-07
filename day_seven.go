package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type HandType int

const (
	high_card HandType = iota + 1
	one_pair
	two_pair
	three_of_a_kind
	full_house
	four_of_a_kind
	five_of_a_kind
)

type Card int

const (
	jack Card = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	queen
	king
	ace
)

func main() {
	hands := ReadFile("inputs/7.txt")
	// log.SetLevel(log.DebugLevel)

	sort.SliceStable(hands, func(i, j int) bool {
		first := hands[i]
		second := hands[j]

		first_cards := conv_to_enum(strings.Fields(first)[0])
		second_cards := conv_to_enum(strings.Fields(second)[0])

		first_hand_type := getHandType(first_cards)
		second_hand_type := getHandType(second_cards)

		if first_hand_type == second_hand_type {
			for i := 0; i < len(first_cards); i++ {
				if first_cards[i] == second_cards[i] {
					continue
				}
				return first_cards[i] < second_cards[i]
			}
		}
		return first_hand_type < second_hand_type
	})

	total := 0

	for i, hand := range hands {
		score, _ := strconv.Atoi(strings.Fields(hand)[1])

		total += (score * (i + 1))
	}
	log.Debug(hands)

	log.Info(total)
}

func getHandType(original_hand []Card) HandType {
	hand := make([]Card, len(original_hand))
	copy(hand, original_hand)

	// count up the number of jacks, and create a new slice without them
	j_count := 0
	var hand_without_j []Card
	for _, card := range hand {
		if card == jack {
			j_count += 1
		} else {
			hand_without_j = append(hand_without_j, card)
		}
	}

	// count instances of the cards
	card_counts := make(map[Card]int)
	for _, card := range hand_without_j {
		card_counts[card] += 1
	}

	// find the most common card, breaking any ties by taking the higher of the two
	most_common_card := jack
	counter := -1
	for next_card, value := range card_counts {
		if value > counter {
			most_common_card = next_card
			counter = value
		} else if value == counter {
			if next_card > most_common_card {
				most_common_card = next_card
				counter = value
			}
		}
	}

	// replace jacks with the new cards to make the best hand possible
	rebuilt_hand := make([]Card, len(hand_without_j))
	copy(rebuilt_hand, hand_without_j)

	for i := 0; i < j_count; i++ {
		rebuilt_hand = append(rebuilt_hand, most_common_card)
	}

	// Sort the hand
	sort.SliceStable(rebuilt_hand, func(i, j int) bool {
		return rebuilt_hand[i] < rebuilt_hand[j]
	})
	log.Debug(rebuilt_hand)

	var matches []int

	// count up the matches to decide how good the hand is
	// n.b. now that we're in part two, I probably could have reused
	// the map I've built above to do this but at least I know this method works
	match_counter := 0
	for i, card := range rebuilt_hand {
		if i+1 < len(rebuilt_hand) && card == rebuilt_hand[i+1] {
			match_counter += 1
		} else if match_counter > 0 {
			matches = append(matches, match_counter)
			match_counter = 0
		}
	}

	sort.Ints(matches)

	if len(matches) == 1 {
		switch matches[0] {
		case 4:
			return five_of_a_kind
		case 3:
			return four_of_a_kind
		case 2:
			return three_of_a_kind
		case 1:
			return one_pair
		}
	} else if len(matches) == 2 {
		if matches[0] == matches[1] {
			return two_pair
		}
		return full_house
	}

	return high_card
}

func conv_to_enum(hand string) (hand_enum []Card) {
	for _, card := range hand {
		hand_enum = append(hand_enum, get_card(card))
	}
	return
}

func get_card(card rune) (card_type Card) {
	switch string(card) {
	case "2":
		return two
	case "3":
		return three
	case "4":
		return four
	case "5":
		return five
	case "6":
		return six
	case "7":
		return seven
	case "8":
		return eight
	case "9":
		return nine
	case "T":
		return ten
	case "J":
		return jack
	case "Q":
		return queen
	case "K":
		return king
	case "A":
		return ace
	default:
		log.Fatalf("Unknown card type %s", string(card))
	}

	return
}
