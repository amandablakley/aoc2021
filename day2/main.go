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

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(input []byte) int {
	inputString := strings.Split(string(input), "\n")
	horizontalPos := 0
	depth := 0

	for _, line := range inputString {
		if line == "" {
			continue
		}

		newString := strings.Split(line, " ")
		direction := newString[0]
		x, _ := strconv.Atoi(newString[1])

		switch direction {
		case "forward":
			horizontalPos += x
		case "up":
			depth -= x
		case "down":
			depth += x
		default:
			//do nothing
		}
	}

	return horizontalPos * depth
}

func partTwo(input []byte) int {
	inputString := strings.Split(string(input), "\n")
	horizontalPos := 0
	depth := 0
	aim := 0

	for _, line := range inputString {
		if line == "" {
			continue
		}

		newString := strings.Split(line, " ")
		direction := newString[0]
		x, _ := strconv.Atoi(newString[1])

		switch direction {
		case "forward":
			horizontalPos += x
			depth += aim * x
		case "up":
			aim -= x
		case "down":
			aim += x
		default:
			// do nothing
		}

	}

	return horizontalPos * depth
}

func getInput() ([]byte, error) {
	return ioutil.ReadFile("./input.input")
}
