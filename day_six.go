package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	file_lines := ReadFile("inputs/6.txt")

	time := strings.Join(strings.Fields(strings.Split(file_lines[0], ":")[1]), "")
	distance := strings.Join(strings.Fields(strings.Split(file_lines[1], ":")[1]), "")

	total := 0

	race_time, _ := strconv.Atoi(time)
	race_distance, _ := strconv.Atoi(distance)
	for button_time := 0; button_time <= race_time; button_time++ {
		remaining_time := race_time - button_time

		if (button_time * remaining_time) > race_distance {
			total += 1
		}
	}

	log.Infof("Total: %d", total)
}
