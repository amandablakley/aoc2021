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

}

func getInput() ([]byte, error) {
	return ioutil.ReadFile("./input.input")
}
