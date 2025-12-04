package main

import (
	"fmt"
	"io"
	"os"
	h "simon-wg/AdventOfCode/util"
)

func main() {
	input, _ := io.ReadAll(os.Stdin)
	data := parseInput(string(input))
	part := h.GetPart()
	fmt.Println(solve(data, part))
}

func parseInput(input string) []int8 {
	return make([]int8, 0)
}

func solve(data []int8, part int) int {
	return 0
}
