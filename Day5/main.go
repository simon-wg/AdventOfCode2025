package main

import (
	"fmt"
	"io"
	"os"
	h "simon-wg/AdventOfCode/util"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	ranges, ingredients := parseInput(string(input))
	part := h.GetPart()
	fmt.Println(solve(ranges, ingredients, part))
}

func parseInput(input string) ([]Interval, []int) {
	split := strings.Split(strings.TrimSpace(input), "\n")
	middle := slices.Index(split, "")
	if middle == -1 {
		middle = len(split)
	}

	rangesStr := split[:middle]
	var ranges []Interval
	for _, s := range rangesStr {
		l, r := realRange(s)
		ranges = append(ranges, Interval{Start: l, End: r})
	}
	var ingredients []int
	if middle < len(split)-1 {
		ingredientsStr := split[middle+1:]
		ingredients = make([]int, len(ingredientsStr))
		for i, e := range ingredientsStr {
			parsed, _ := strconv.Atoi(e)
			ingredients[i] = int(parsed)
		}
	}

	return ranges, ingredients
}

func solve(ranges []Interval, ingredients []int, part int) int {
	if part == 1 {
		sol := 0
		for _, ingredient := range ingredients {
			for _, r := range ranges {
				if ingredient >= r.Start && ingredient <= r.End {
					sol++
					break
				}
			}
		}
		return sol
	} else {
		slices.SortFunc(ranges, func(a, b Interval) int {
			return a.Start - b.Start
		})

		merged := []Interval{ranges[0]}

		for _, r := range ranges[1:] {
			last := &merged[len(merged)-1]
			if r.Start <= last.End+1 {
				if r.End > last.End {
					last.End = r.End
				}
			} else {
				merged = append(merged, r)
			}
		}

		totalFresh := 0
		for _, r := range merged {
			totalFresh += (r.End - r.Start + 1)
		}
		return totalFresh
	}
}

func realRange(s string) (int, int) {
	split := strings.Split(s, "-")
	left, _ := strconv.Atoi(split[0])
	right, _ := strconv.Atoi(split[1])
	return int(left), int(right)
}
