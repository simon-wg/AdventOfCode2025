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
	return strings.SplitSeq(strings.TrimSpace(string(input)), ",")
}

func solve(data iter.Seq[string], part int) int {
	sol := 0
	for s := range data {
		splits := strings.Split(s, "-")
		start, _ := strconv.Atoi(splits[0])
		end, _ := strconv.Atoi(splits[1])
		for i := start; i <= end; i++ {
			sVal := strconv.Itoa(i)
			length := len(sVal)
			pLen := 1
			if part == 1 {
				pLen = length / 2
			}
			for ; pLen <= length/2; pLen++ {
				if length%pLen == 0 {
					if strings.Repeat(sVal[:pLen], length/pLen) == sVal {
						sol += i
						break
					}
				}
			}
		}
	}
	return sol
}
