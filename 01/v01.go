package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("01/01_real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	max := 0
	total := 0
	var calories []int
	for number, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			intval, _ := strconv.Atoi(line)
			total += intval
		} else {
			calories = append(calories, total)
			total = 0
		}
		fmt.Println("as of", number, "total is", total)
		if total > max {
			max = total
		}
	}
	fmt.Println("Part 1:", max)
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println("Part 2:", calories[0]+calories[1]+calories[2])

}
