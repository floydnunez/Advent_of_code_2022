package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("06/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	line := strings.Split(lines[0], "")
	length := len(line)
	result1 := part1(length, line)
	println("\nPart 1:", result1) //1757
	result2 := part2(length, line)
	println("\nPart 2:", result2) //2950
}

func part2(length int, line []string) int {
	for pos := 14; pos < length; pos++ {
		//sub := strings.Join(line[pos-14:pos], "")
		isSignal := checkSignal2(line[pos-14 : pos])
		//println(isSignal, pos, sub)
		if isSignal {
			return pos
		}
	}
	return -1
}

func checkSignal2(sub []string) bool {
	m := make(map[string]bool)
	for _, elem := range sub {
		m[elem] = true
	}
	return len(m) == 14
}

func part1(length int, line []string) int {
	for pos := 4; pos < length; pos++ {
		sub := strings.Join(line[pos-4:pos], "")
		isSignal := checkSignal(sub)
		//println(isSignal, pos, sub)
		if isSignal {
			return pos
		}
	}
	return -1
}

func checkSignal(line string) bool {
	//println(line[0] != line[1], line[0] != line[2], line[0] != line[3],
	//	line[1] != line[2], line[1] != line[3], line[2] != line[3])
	return (line[0] != line[1]) && (line[0] != line[2]) && (line[0] != line[3]) &&
		(line[1] != line[2]) && (line[1] != line[3]) && (line[2] != line[3])
}
