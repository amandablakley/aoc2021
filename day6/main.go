package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func main() {
	input, err := ioutil.ReadFile("./input.input")

	if err != nil {
		log.Println(err)
	}

	lines := strings.Split(string(input), ",")

	inputMap := make([]int64, 9)
	for i := 0; i <= 8; i++ {
		input[i] = 0
	}
	for _, value := range lines {
		val, _ := strconv.ParseInt(value, 10, 16)
		inputMap[uint(val)] += 1
	}

	totalDayOne, totalDayTwo := days(inputMap, 80, 256)

	fmt.Println("Part One: ", totalDayOne)
	fmt.Println("Part Two: ", totalDayTwo)

}

func oneDay(input []int64) []int64 {
	total := make([]int64, 9)
	for _, i := range [7]uint{0, 1, 2, 3, 4, 5, 7} {
		total[i] = input[i+1]
	}
	total[6] = input[0] + input[7]
	total[8] = input[0]
	return total
}

func days(input []int64, day1, day2 uint16) (int64, int64) {
	i := uint16(0)

	for ; i < day1; i++ {
		input = oneDay(input)
	}

	totalDayOne := int64(0)
	for _, value := range input {
		totalDayOne += value
	}

	// Part 2
	for ; i < day2; i++ {
		input = oneDay(input)
	}

	totalDayTwo := int64(0)
	for _, value := range input {
		totalDayTwo += value
	}

	return totalDayOne, totalDayTwo
}
