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
	dat, err := os.ReadFile("10/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	totalP1 := 0
	x := 1
	cycle := 0
	var interesting []int
	var screen [240]string
	screen = cleanScreen(screen)
	for _, line := range lines {
		printScreen(screen, cycle)
		if line == "noop" {
			cycle += 1
			screen = draw(cycle, screen, x)
			interesting = checkCycle(interesting, cycle, x)
		} else {
			cycle += 1
			println("pre:", cycle, line, x)
			screen = draw(cycle, screen, x)
			interesting = checkCycle(interesting, cycle, x)
			println(cycle, line, x)
			parts := strings.Split(line, " ")
			amount, _ := strconv.Atoi(parts[1])
			cycle += 1
			//this has to happen BEFORE the add. The addition only happens after the cycle ends
			screen = draw(cycle, screen, x)
			interesting = checkCycle(interesting, cycle, x)
			x += amount
		}
		println(cycle, line, x)
	}
	fmt.Println(interesting)
	for _, val := range interesting {
		totalP1 += val
	}
	fmt.Println("Part 1:", totalP1) //14320
	//part 2: PCPBKAPJ
}

func printScreen(screen [240]string, cycle int) {
	if (cycle+1)%240 == 0 {
		for num, val := range screen {
			print(val)
			if (num+1)%40 == 0 {
				println("")
			}
		}
		screen = cleanScreen(screen)
	}
	return
}

func draw(cycle int, screen [240]string, x int) [240]string {
	pos := cycle % 40
	row := cycle / 40
	fmt.Println("draw: ", pos, x)
	if pos >= x && pos < x+3 {
		screen[pos+40*row-1] = "#"
	}
	return screen
}

func cleanScreen(screen [240]string) [240]string {
	for num := range screen {
		screen[num] = " "
	}
	return screen
}

func checkCycle(interesting []int, cycle int, x int) []int {
	if cycle == 20 || (cycle-20)%40 == 0 {
		println("interesting!", cycle, x)
		interesting = append(interesting, x*cycle)
	}
	return interesting
}
