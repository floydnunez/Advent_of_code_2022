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

func main() {
	dat, err := os.ReadFile("00/example.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	max := 0
	for _, line := range lines {
		println(line)
	}
	fmt.Println("Part 1:", max)
	fmt.Println("Part 2:", max)

}
