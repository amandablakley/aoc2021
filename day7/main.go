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

	var pos []int
	max := 0

	for _, s := range strings.Split(string(input), ",") {
		i, err := strconv.Atoi(s)

		if err != nil {
			log.Println(err)
			return
		}

		if i > max {
			max = i
		}

		pos = append(pos, i)
	}

	distance := make([]int, max+1)

	for i := 0; i <= max; i++ {
		for _, n := range pos {
			distance[i] += fuelCost(abs(n - i))
		}
	}

	min := 100000000
	for _, d := range distance {
		if d < min {
			min = d
		}
	}

	fmt.Println(min)
}

func fuelCost(num int) int {
	return (num * (num + 1)) / 2
}

func abs(in int) int {
	if in < 0 {
		return in * -1
	}

	return in
}

func getInput() ([]byte, error) {
	return ioutil.ReadFile("./input.input")
}
