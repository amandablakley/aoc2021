package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.input")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(input), "\n")

	// fmt.Println(partOne(lines))
	partTwo(lines)
}

func partOne(input string) int64 {
	gamma := ""
	epsilon := ""
	inputNumbers := strings.Split(input, "\n")

	i := 0

	for i < len(inputNumbers[0]) {
		zeroes := 0
		ones := 0

		for _, number := range inputNumbers {
			if number[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}

		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}

		i++
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return g * e
}

func partTwo(lines []string) {
	i := 0
	nl := make([]string, len(lines))
	copy(nl, lines)

	for len(nl) > 1 {
		nl = filter(nl, i, false)
		i++
	}

	fmt.Println(nl)

	oxygen, _ := strconv.ParseInt(nl[0], 2, 64)

	i = 0
	co2Lines := make([]string, len(lines))
	copy(co2Lines, lines)

	for len(co2Lines) > 1 {
		co2Lines = filter(co2Lines, i, true)
		i++
	}

	co2, _ := strconv.ParseInt(co2Lines[0], 2, 64)

	fmt.Println(oxygen * co2)
}

func filter(in []string, pos int, lowest bool) []string {
	ones, zeros := 0, 0
	for _, l := range in {
		if l[pos] == '1' {
			ones++
		} else {
			zeros++
		}
	}

	var bit byte = '1'
	if (lowest && ones >= zeros) || (!lowest && ones < zeros) {
		bit = '0'
	}

	count := 0
	for i := 0; i < len(in); i++ {
		if in[i][pos] == bit {
			in[count] = in[i]
			count++
		}
	}

	return in[:count]
}
