package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("inputs/4.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// total := 0

	var cards []string
	var queue []int

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		substrings := strings.Split(line, ":")
		cards = append(cards, substrings[1])
		queue = append(queue, i)
		i++
	}

	total_cards := 0

	for len(queue) > 0 {
		next_id := queue[0]
		queue = queue[1:]
		total_cards += 1

		game_row := cards[next_id]

		game := strings.Split(game_row, "|")

		winning_nums := strings.Fields(game[0])
		card_nums := strings.Fields(game[1])

		match_counter := 0

		for _, num := range card_nums {
			for _, winner := range winning_nums {
				if num == winner {
					match_counter += 1
				}
			}
		}

		for i := 1; i <= match_counter; i++ {
			queue = append(queue, next_id+i)
		}

	}

	println(total_cards)

}
