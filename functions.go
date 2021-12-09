package functions

import "strconv"

func min(minNum []int) int {
	min := minNum[0]

	for i := 0; i < len(minNum); i++ {
		if minNum[i] < min {
			min = minNum[i]
		}
	}

	return min
}

func max(maxNum []int) int {
	max := 0

	for i := 0; i < len(maxNum); i++ {
		if maxNum[i] > max {
			max = maxNum[i]
		}
	}

	return max
}

func convertInt(input []string) []int {
	number := []int{}

	for _, n := range input {
		i, _ := strconv.Atoi(n)
		number = append(number, i)
	}

	return number
}

func abs(in int) int {
	if in < 0 {
		return in * -1
	}

	return in
}
