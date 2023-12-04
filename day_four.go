package main

import (
	"bufio"
	"fmt"
	"math"
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

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		substrings := strings.Split(line, ":")

		game := strings.Split(substrings[1], "|")

		fmt.Println(substrings[0])

		winning_nums := strings.Split(game[0], " ")
		card_nums := strings.Split(game[1], " ")

		match_counter := 0

		for _, num := range card_nums {
			for _, winner := range winning_nums {
				if num != "" && winner != "" && num == winner {
					match_counter += 1
				}
			}
		}

		if match_counter == 1 {
			total += 1
		} else {

			additional_points := match_counter - 1
			base := 1
			if additional_points < 0 {
				additional_points = 0
				base = 0
			}
			match_points := base * (int(math.Pow(2, float64(additional_points))))

			total += match_points
		}
	}

	fmt.Println(total)
}
