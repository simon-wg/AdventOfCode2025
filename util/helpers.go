package helpers

import (
	"fmt"
	"os"
)

func GetPart() int {
	if len(os.Args) < 2 {
		fmt.Println("Run using `go run main.go <part> < input.txt`")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1":
		return 1
	case "2":
		return 2
	default:
		fmt.Println("Invalid part, only 1 and 2 exist. Defaulting to part 2")
		return 2
	}
}

func FloorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}
	return (a - b + 1) / b
}

func Mod(a, b int) int {
	return (a%b + b) % b
}
