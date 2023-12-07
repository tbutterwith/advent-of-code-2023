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
	two Card = iota + 1
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
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

		log.Debugf("%s (%v), %s(%v)", first, first_hand_type, second, second_hand_type)
		if first_hand_type == second_hand_type {
			for i := 0; i < len(first_cards); i++ {

				log.Debugf("%v:%v", first_cards[i], second_cards[i])
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

func getHandType(hand []Card) HandType {
	hand_clone := make([]Card, len(hand))
	copy(hand_clone, hand)

	sort.SliceStable(hand_clone, func(i, j int) bool {
		return hand_clone[i] < hand_clone[j]
	})

	var matches []int

	match_counter := 0
	for i, card := range hand_clone {
		if i+1 < len(hand_clone) && card == hand_clone[i+1] {
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
