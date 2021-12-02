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
	intDepths := convertInt(seaDepths)

	for i := 0; i < len(intDepths)-1; i++ {
		if i == 0 {
			continue
		}

		next := seaDepths[i]
		previous := seaDepths[i-1]
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

	for i := 2; i < len(intDepths)-1; i++ {
		if i == 0 || intDepths[i] == 0 {
			continue
		}

		currentWindow := intDepths[i] + intDepths[i-1] + intDepths[i-2]
		nextWindow := intDepths[i] + intDepths[i+1] + intDepths[i-1]

		if currentWindow < nextWindow {
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

func getInput() ([]byte, error) {
	return ioutil.ReadFile("./input.input")
}
