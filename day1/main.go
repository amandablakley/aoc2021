package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := getInput()

	if err != nil {
		log.Println(err)
		return
	}

	partOne(input)
	partTwo(input)
}

func partOne(input []byte) {
	increment := 0
	seaDepths := strings.Split(string(input), "\n")

	for i := 0; i < len(seaDepths); i++ {
		if i == 0 || seaDepths[i] == "" {
			continue
		}

		next, _ := strconv.Atoi(seaDepths[i])
		previous, _ := strconv.Atoi(seaDepths[i-1])
		incremented := next > previous

		if incremented {
			increment++
		}
	}

	fmt.Println("Part One:", increment, "increments")
}

func partTwo(input []byte) {
	larger := 0
	seaDepths := strings.Split(string(input), "\n")
	intDepths := convertInt(seaDepths)

	for i := 0; i < len(intDepths); i++ {
		if i == 0 || i+3 >= len(intDepths) || intDepths[i] == 0 {
			continue
		}

		windowSum := intDepths[i] + intDepths[i+1] + intDepths[i+2]
		nextWindow := intDepths[i+1] + intDepths[i+2] + intDepths[i+3]

		if windowSum < nextWindow {
			larger++
		}

	}

	fmt.Println("Part Two:", larger, "larger")

}

func convertInt(input []string) []int {
	number := []int{}

	for _, n := range input {
		i, _ := strconv.Atoi(n)
		number = append(number, i)
	}

	return number
}

func getSum(input []int) int {
	sum := 0

	for _, v := range input {
		sum += v
	}

	return sum
}

func getInput() ([]byte, error) {
	return ioutil.ReadFile("./input.input")
}
