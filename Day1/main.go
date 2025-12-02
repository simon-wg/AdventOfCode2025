package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var curr = 50
var passwd = 0

func main() {
	input, _ := os.ReadFile("Day1/in.txt")
	operations := strings.SplitSeq(string(input), "\n")
	for op := range operations {
		if len(op) == 0 {
			break
		}
		turn(op)
	}
	fmt.Println(passwd)
}

func turn(op string) {
	direction := op[0]
	amount, _ := strconv.Atoi(op[1:])
	switch direction {
	case 'R':
		turnRight(amount)
	case 'L':
		turnLeft(amount)
	}
}

func turnRightPart1(amount int) {
	curr = mod(curr+amount, 100)
	if curr == 0 {
		passwd++
	}
}

func turnLeftPart1(amount int) {
	curr = mod(curr-amount, 100)
	if curr == 0 {
		passwd++
	}
}

func turnRight(amount int) {
	passwd += (curr + amount) / 100
	curr = mod(curr+amount, 100)
}

func turnLeft(amount int) {
	term1 := floorDiv(curr-1, 100)
	term2 := floorDiv(curr-amount-1, 100)
	passwd += term1 - term2
	curr = mod(curr-amount, 100)
}

func floorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return (a - b + 1) / b
}

func mod(a, b int) int {
	return (a%b + b) % b
}
