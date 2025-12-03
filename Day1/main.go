package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	h "simon-wg/AdventOfCode/util"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	data := parseInput(input)
	part := h.GetPart()
	fmt.Println(solve(data, part))
}

func parseInput(input []byte) iter.Seq[string] {
	return strings.SplitSeq(string(input), "\n")
}

func solve(data iter.Seq[string], part int) int {
	var curr = 50
	var passwd = 0

	for op := range data {
		if len(op) == 0 {
			break
		}
		if part == 1 {
			part1(op, &curr, &passwd)
			continue
		}
		part2(op, &curr, &passwd)
	}
	return passwd
}

func part1(op string, curr, passwd *int) {
	direction := op[0]
	amount, _ := strconv.Atoi(op[1:])
	switch direction {
	case 'R':
		*curr = h.Mod(*curr+amount, 100)
	case 'L':
		*curr = h.Mod(*curr-amount, 100)
	}
	if *curr == 0 {
		*passwd++
	}
}

func part2(op string, curr, passwd *int) {
	direction := op[0]
	amount, _ := strconv.Atoi(op[1:])
	switch direction {
	case 'R':
		*passwd += (*curr + amount) / 100
		*curr = h.Mod(*curr+amount, 100)
	case 'L':
		term1 := h.FloorDiv(*curr-1, 100)
		term2 := h.FloorDiv(*curr-amount-1, 100)
		*passwd += term1 - term2
		*curr = h.Mod(*curr-amount, 100)
	}
}
