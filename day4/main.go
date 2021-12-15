package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type num struct {
	num     int16
	checked bool
}

func main() {
	input, err := ioutil.ReadFile("./input.input")

	if err != nil {
		log.Println(err)
	}

	lines := strings.Split(string(input), "\n\n")

	nums := strings.Split(lines[0], ",")
	var numbers []int16

	for _, v := range nums {
		val, _ := strconv.ParseInt(v, 10, 16)
		numbers = append(numbers, int16(val))
	}

	//Read the bingos
	bingosAux := lines[1:]
	var bingos [][][]num
	for _, v := range bingosAux {
		bingosAux2 := strings.Split(v, "\n")

		// Read each board
		var bingos2 [][]num
		for _, v2 := range bingosAux2 {
			if v2[0] == ' ' {
				v2 = v2[1:]
			}

			bingosAux3 := strings.Split(strings.Replace(v2, "  ", " ", -1), " ")
			var bingos3 []num

			for _, v3 := range bingosAux3 {
				val, _ := strconv.ParseInt(v3, 10, 16)
				bingos3 = append(bingos3, num{num: int16(val), checked: false})
			}

			bingos2 = append(bingos2, bingos3)
		}
		bingos = append(bingos, bingos2)
	}

	part1 := part1(numbers, bingos)
	part2 := part2(numbers, bingos)

	fmt.Println("part 1", part1, "part 2", part2)
}

func part1(nums []int16, bingos [][][]num) int64 {
	for _, num := range nums {
		bingos = updateBingo(num, bingos)
		for _, b := range bingos {
			if checkBingoBoard(b) {
				return int64(getUnMarked(b)) * int64(num)
			}
		}
	}

	return 0
}

func updateBingo(num int16, bingos [][][]num) [][][]num {
	for n := 0; n < len(bingos); n++ {
		for i := 0; i < len(bingos[0]); i++ {
			for j := 0; j < len(bingos[0][0]); j++ {
				if num == bingos[n][i][j].num {
					bingos[n][i][j].checked = true
				}
			}
		}
	}

	return bingos
}

func checkBingoBoard(bingo [][]num) bool {
	for i := 0; i < len(bingo); i++ {
		line := true
		for j := 0; j < len(bingo[0]); j++ {
			line = line && bingo[i][j].checked
		}

		if line {
			return true
		}
	}

	for j := 0; j < len(bingo); j++ {
		column := true
		for k := 0; k < len(bingo[0]); k++ {
			column = column && bingo[k][j].checked
		}

		if column {
			return true
		}
	}

	return false
}

func getUnMarked(bingo [][]num) int16 {
	sum := int16(0)

	for _, v := range bingo {
		for _, val := range v {
			if !val.checked {
				sum += val.num
			}
		}
	}

	return sum
}

func part2(nums []int16, bingos [][][]num) int64 {
	notWon := makeRange(0, len(bingos))
	var last int

	for _, num := range nums {
		bingos = updateBingo(num, bingos)

		for n, b := range bingos {
			if checkBingoBoard(b) {
				notWon = removeBoard(notWon, n)
			}

			if len(notWon) == 1 {
				last = notWon[0]
			} else if len(notWon) == 0 {
				return int64(getUnMarked(bingos[last])) * int64(num)
			}
		}
	}

	return 0
}

func makeRange(min, max int) []int {
	m := make([]int, max-min)
	for i := range m {
		m[i] = min + i
	}

	return m
}

func removeBoard(board []int, item int) []int {
	for i, j := range board {
		if j == item {
			return append(board[:i], board[i+1:]...)
		}
	}

	return board
}

func getInput() ([]byte, error) {
	return ioutil.ReadFile("./input.input")
}
