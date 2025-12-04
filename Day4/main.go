package main

import (
	"fmt"
	"io"
	"os"
	h "simon-wg/AdventOfCode/util"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	data := parseInput(string(input))
	part := h.GetPart()
	fmt.Println(solve(data, part))
}

func parseInput(input string) [][]int {
	rowCount := 0
	colCount := 0
	rows := strings.Split(strings.TrimSpace(input), "\n")
	for _, row := range rows {
		rowCount++
		colCount = len(row)
	}
	data := make([][]int, rowCount)
	for y, row := range rows {
		data[y] = make([]int, colCount)
		for x := range len(row) {
			if row[x] == '@' {
				data[y][x] = 1
				continue
			}
		}
	}
	return data
}

func solve(data [][]int, part int) int {
	sol := 0
	done := true
	for {
		for y := range len(data) {
			for x := range len(data[0]) {
				if check(x, y, data) < 4 {
					if data[y][x] == 1 {
						sol++
						done = part == 1
					}
					if part == 2 {
						data[y][x] = 0
					}
				}
			}
		}
		if done {
			break
		}
		done = true
	}
	return sol
}

func check(x, y int, data [][]int) int {
	total := 0
	if x > 0 {
		if y > 0 {
			total += data[y-1][x-1]
		}
		if y < len(data[0])-1 {
			total += data[y+1][x-1]
		}
		total += data[y][x-1]
	}
	if y > 0 {
		total += data[y-1][x]
	}
	if x < len(data[0])-1 {
		if y < len(data)-1 {
			total += data[y+1][x+1]
		}
		if y > 0 {
			total += data[y-1][x+1]
		}
		total += data[y][x+1]
	}
	if y < len(data)-1 {
		total += data[y+1][x]
	}
	return total
}
