package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sol = 0

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Run using `go run ./Day3 <part>`")
		fmt.Println("E.g. `go run ./Day3 1`")
		os.Exit(1)
	}
	part := os.Args[1]
	nDigits := 0
	switch part {
	case "1":
		nDigits = 2
	case "2":
		nDigits = 12
	default:
		fmt.Println("Only part 1 and 2 exist")
		fmt.Println("Run using `go run ./Day3 <part>`")
		fmt.Println("E.g. `go run ./Day3 1`")
		os.Exit(1)
	}

	input, _ := os.ReadFile("Day3/in.txt")
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
	for _, row := range data {
		var digits []int = make([]int, nDigits)
		lastPos := -1
		for i := nDigits; i > 0; i-- {
			digits[i-1], lastPos = findDigit(row, lastPos+1, i-1)
		}
		for i := nDigits; i > 0; i-- {
			digit := digits[i-1] * int(math.Pow10(i-1))
			sol += digit
			fmt.Println(digit)
		}
	}
	fmt.Println(sol)
}

func findDigit(row []int, pos int, n int) (int, int) {
	largest := 0
	for i := pos; i < len(row)-n; i++ {
		if row[i] > largest {
			largest = row[i]
			pos = i
		}
	}
	return largest, pos
}
