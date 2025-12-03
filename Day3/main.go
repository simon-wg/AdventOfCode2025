package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"simon-wg/AdventOfCode/util"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	data := parseInput(input)
	part := helpers.GetPart()
	fmt.Println(solve(data, part))
}

func parseInput(input []byte) [][]int {
	rowCount := 0
	colCount := 0
	rows := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, row := range rows {
		rowCount++
		colCount = len(string(row))
	}
	var data [][]int = make([][]int, rowCount)
	for i := range data {
		data[i] = make([]int, colCount)
	}
	for y, row := range rows {
		for x := range colCount {
			n, _ := strconv.Atoi(string(row[x]))
			data[y][x] = n
		}
	}
	return data
}

func solve(data [][]int, part int) int {
	sol := 0
	nDigits := 12
	if part == 1 {
		nDigits = 2
	}
	for _, row := range data {
		var digits []int = make([]int, nDigits)
		lastPos := -1
		for i := nDigits; i > 0; i-- {
			digits[i-1], lastPos = findDigit(row, lastPos+1, i-1)
		}
		for i := nDigits; i > 0; i-- {
			digit := digits[i-1] * int(math.Pow10(i-1))
			sol += digit
		}
	}
	return sol
}

func findDigit(row []int, pos int, n int) (int, int) {
	largest := 0
	for i := pos; i < len(row)-n; i++ {
		if row[i] == 9 {
			return 9, i
		}
		if row[i] > largest {
			largest = row[i]
			pos = i
		}
	}
	return largest, pos
}
