package main

import (
	"fmt"
	"io"
	"os"
	"simon-wg/AdventOfCode/util"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	data := parseInput(input)
	part := helpers.GetPart()
	fmt.Println(solve(data, part))
}

func parseInput(input []byte) [][]int {
	return make([][]int, 0)
}

func solve(data [][]int, part int) int {
	return 0
}
