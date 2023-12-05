package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type dest_tupe struct {
	dest int
	len  int
}

var seed_to_soil map[int]dest_tupe
var soil_to_fertiliser map[int]dest_tupe
var fertiliser_to_water map[int]dest_tupe
var water_to_light map[int]dest_tupe
var light_to_temp map[int]dest_tupe
var temp_to_humidity map[int]dest_tupe
var humidity_to_loc map[int]dest_tupe

func main() {
	file_lines := ReadFile("inputs/5.txt")

	seeds := strings.Fields(strings.Split(file_lines[0], ":")[1])
	seed_to_soil = make(map[int]dest_tupe)
	soil_to_fertiliser = make(map[int]dest_tupe)
	fertiliser_to_water = make(map[int]dest_tupe)
	water_to_light = make(map[int]dest_tupe)
	light_to_temp = make(map[int]dest_tupe)
	temp_to_humidity = make(map[int]dest_tupe)
	humidity_to_loc = make(map[int]dest_tupe)

	map_ids := []string{"soil", "fert", "water", "light", "temp", "humidity", "location"}
	maps := []map[int]dest_tupe{seed_to_soil, soil_to_fertiliser, fertiliser_to_water, water_to_light, light_to_temp, temp_to_humidity, humidity_to_loc}

	map_iterator := 0

	for _, line := range file_lines[3:] {
		if len(line) == 0 {
			continue
		}
		first_char := rune(line[0])
		if unicode.IsLetter(first_char) {
			map_iterator += 1
		} else if unicode.IsDigit(first_char) {
			current_map := maps[map_iterator]

			nums := strings.Fields(line)
			dest, _ := strconv.Atoi(nums[0])
			source, _ := strconv.Atoi(nums[1])
			ranges, _ := strconv.Atoi(nums[2])

			current_map[source] = dest_tupe{dest, ranges}
		}
	}

	lowest_seed_loc := -1
	for _, seed_str := range seeds {
		seed_num, _ := strconv.Atoi(seed_str)

		id := seed_num
		fmt.Printf("Id is seed: %d\n", id)

		for i, cur_map := range maps {
			id = search(cur_map, id)
			fmt.Printf("Id is %s: %d\n", map_ids[i], id)
		}

		fmt.Printf("Location is %d\n", id)
		fmt.Printf("\n\n\n")
		if id < lowest_seed_loc || lowest_seed_loc == -1 {
			lowest_seed_loc = id
		}
	}

	println(lowest_seed_loc)
}

func search(lookup map[int]dest_tupe, num int) (id int) {
	id = num

	fmt.Printf("looking for id %d\n", id)
	for source, dest_pair := range lookup {
		length := dest_pair.len

		if id >= source && id < source+length {
			// it's a match

			diff := id - source
			fmt.Printf("id is %d, source is %d, length is %d, diff is %d\n", id, source, length, diff)
			id = dest_pair.dest + diff

			return
		}
	}

	return
}
