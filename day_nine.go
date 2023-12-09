package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	lines := ReadFile("inputs/9.txt")
	log.SetLevel(log.DebugLevel)

	var sequences map[int][][]int
	sequences = make(map[int][][]int)

	for i, line := range lines {
		sequence := strings.Fields(line)

		var history [][]int
		var int_sequence []int

		for _, num_str := range sequence {
			num, _ := strconv.Atoi(num_str)
			int_sequence = append(int_sequence, num)
		}

		history = append(history, int_sequence)
		sequences[i] = history
	}

	total := 0
	for _, sequence := range sequences {
		new_sequence := calc_history(sequence)
		num := infer_new_nums_forward(new_sequence)

		total += num
	}

	log.Info(total)
}

func calc_history(sequence [][]int) (new_sequence [][]int) {
	var queue [][]int
	queue = append(queue, sequence[0])

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		diffs := get_diffs(next)

		sequence = append(sequence, diffs)

		all_zero := true
		for _, x := range diffs {
			if x != 0 {
				all_zero = false
			}
		}
		if !all_zero {
			queue = append(queue, diffs)
		}
	}

	log.Debugf("Sequence with history: %v", sequence)
	new_sequence = sequence
	return
}

func get_diffs(row []int) (diffs []int) {
	for i, num := range row {
		if i < len(row)-1 {
			next_num := row[i+1]

			diff := next_num - num
			// abs_diff := math.Abs(float64(diff))
			// diffs = append(diffs, int(abs_diff))
			diffs = append(diffs, diff)

		}
	}

	return
}

func infer_new_nums_forward(sequence [][]int) (final_val int) {
	is_first_row := true

	for i := len(sequence) - 2; i >= 0; i-- {
		next_num := calc_next_num(sequence[i], sequence[i+1], is_first_row)
		sequence[i] = append(sequence[i], next_num)
		is_first_row = false
	}

	log.Debug(sequence)
	first_seq := sequence[0]
	final_val = first_seq[len(first_seq)-1]

	log.Debug(final_val)
	return
}

func calc_next_num(row, prev_row []int, diff_is_zero bool) (next_num int) {
	if diff_is_zero {
		next_num = row[len(row)-1]
	}

	diff := prev_row[len(prev_row)-1]
	next_num = row[len(row)-1] + diff
	return
}
