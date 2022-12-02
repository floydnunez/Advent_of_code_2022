package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ownSelectionPoints(line string) int {
	if strings.HasSuffix(line, "X") {
		return 1
	}
	if strings.HasSuffix(line, "Y") {
		return 2
	}
	return 3
}

func matchPoints(m map[string]int, line string) int {
	return m[line]
}

func main() {
	m := make(map[string]int)
	m["A X"] = 3
	m["A Y"] = 6
	m["A Z"] = 0
	m["B X"] = 0
	m["B Y"] = 3
	m["B Z"] = 6
	m["C X"] = 6
	m["C Y"] = 0
	m["C Z"] = 3

	dat, err := os.ReadFile("02/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	total := 0
	for _, line := range lines {
		total += ownSelectionPoints(line) + matchPoints(m, line)
	}
	fmt.Println("Part 1:", total)
	total = 0

	n := make(map[string]string)
	n["A X"] = "Z"
	n["A Y"] = "X"
	n["A Z"] = "Y"
	n["B X"] = "X"
	n["B Y"] = "Y"
	n["B Z"] = "Z"
	n["C X"] = "Y"
	n["C Y"] = "Z"
	n["C Z"] = "X"

	for _, line := range lines {
		total += necessaryOutcomePoints(line) + necessaryChoicePoints(n, line)
	}
	fmt.Println("Part 2:", total)

}

func necessaryChoicePoints(n map[string]string, line string) int {
	return ownSelectionPoints(n[line])
}

func necessaryOutcomePoints(line string) int {
	if strings.HasSuffix(line, "X") {
		return 0
	}
	if strings.HasSuffix(line, "Y") {
		return 3
	}
	return 6
}
