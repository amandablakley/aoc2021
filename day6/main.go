package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	lines := strings.Split(string(input), "\n")

	fmt.Println(lines)
}
