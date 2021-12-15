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

	in := strings.Split(string(input), "\n")
	lines := make([][]Coords, len(in))

	for i := 0; i < len(in); i++ {
		points := strings.Split(in[i], " -> ")
		lines[i] = make([]Coords, len(points))

		for j := 0; j < len(points); j++ {
			point := strings.Split(points[j], ",")

			x, err := strconv.Atoi(point[0])

			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.Atoi(point[1])

			if err != nil {
				log.Fatal(err)
			}

			lines[i][j] = Coords{
				x: x,
				y: y,
			}
		}
	}

	grid := make([][]int, 10000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 10000)
	}

	for i := 0; i < len(lines); i++ {
		point1 := lines[i][0]
		point2 := lines[i][1]

		if point1.x != point2.x && point1.y != point2.y {
			continue
		}

		x := point1.x
		y := point1.y
		grid[x][y]++

		for x != point2.x || y != point2.y {
			if x < point2.x {
				x++
			} else if x > point2.x {
				x--
			}

			if y < point2.y {
				y++
			} else if y > point2.y {
				y--
			}

			grid[x][y]++
		}
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				count++
			}
		}
	}

	// Part 2

	grid = make([][]int, 10000)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 10000)
	}

	for i := 0; i < len(lines); i++ {
		point1 := lines[i][0]
		point2 := lines[i][1]
		x := point1.x
		y := point1.y
		grid[x][y]++

		for x != point2.x || y != point2.y {
			if x < point2.x {
				x++
			} else if x > point2.x {
				x--
			}

			if y < point2.y {
				y++
			} else if y > point2.y {
				y--
			}

			grid[x][y]++
		}
	}

	count2 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				count2++
			}
		}
	}

	fmt.Println("part 1:", count, "part 2:", count2)
}
