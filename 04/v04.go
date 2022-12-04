package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("04/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	totalP1 := 0
	totalP2 := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		first := strings.Split(parts[0], "-")
		secon := strings.Split(parts[1], "-")
		a, _ := strconv.Atoi(first[0])
		b, _ := strconv.Atoi(first[1])
		c, _ := strconv.Atoi(secon[0])
		d, _ := strconv.Atoi(secon[1])
		if (a <= c && b >= d) || (c <= a && d >= b) {
			totalP1 += 1
		}
		if (a >= c && a <= d) || (b >= c && b <= d) || (c >= a && c <= b) || (d >= a && d <= b) {
			totalP2 += 1
		}
	}
	fmt.Println("Part 1:", totalP1) //547
	fmt.Println("Part 2:", totalP2) //843
}
