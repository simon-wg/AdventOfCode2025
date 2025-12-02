package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sol = 0

func main() {
	input, _ := os.ReadFile("Day2/in.txt")
	operations := strings.SplitSeq(strings.TrimSpace(string(input)), ",")
	for op := range operations {
		testRange(op)
	}
	fmt.Println(sol)
}

func testRange(s string) {
	splits := strings.Split(s, "-")
	start, _ := strconv.Atoi(splits[0])
	end, _ := strconv.Atoi(splits[1])
	for i := start; i <= end; i++ {
		sVal := strconv.Itoa(i)
		length := len(sVal)
		for pLen := 1; pLen <= length/2; pLen++ {
			if length%pLen == 0 {
				if strings.Repeat(sVal[:pLen], length/pLen) == sVal {
					sol += i
					break
				}
			}
		}
	}
}

