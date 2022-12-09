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
	dat, err := os.ReadFile("08/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	height := len(lines)
	total := 0
	var forest [][]int
	var width int
	for _, line := range lines {
		forest = append(forest, spliceAndParse(line))
		//fmt.Println(row, forest[row])
		width = len(line)
	}
	for row, line := range forest {
		for col, tree := range line {
			total += checkVisibility(forest, col, row, tree, width, height)
		}
	}
	fmt.Println("Part 1:", total) //1717
	part2 := 0
	for row, line := range forest {
		//fmt.Println(line)
		for col, tree := range line {
			l := visIndexLeft(forest, col, row, tree)
			r := visIndexRight(forest, col, row, tree, width)
			u := visIndexUp(forest, col, row, tree)
			d := visIndexDown(forest, col, row, tree, height)
			score := l * r * u * d
			//fmt.Println("for ", col, row, tree, width, height, "l =", l, "r =", r, "u =", u, "d =", d, "score =", score)
			if score >= part2 {
				part2 = score
			}
		}
	}
	fmt.Println("Part 2:", part2) //321975

}

func visIndexLeft(forest [][]int, col int, row int, tree int) int {
	if col == 0 {
		return 0
	}
	count := 0
	for x := col - 1; x >= 0; x-- {
		if forest[row][x] < tree {
			count++
		} else if forest[row][x] >= tree {
			count++
			break
		}
	}
	return count
}

func visIndexRight(forest [][]int, col int, row int, tree int, width int) int {
	if col == width-1 {
		return 0
	}
	count := 0
	for x := col + 1; x < width; x++ {
		if forest[row][x] < tree {
			count++
		} else if forest[row][x] >= tree {
			count++
			break
		}
	}
	return count
}

func visIndexUp(forest [][]int, col int, row int, tree int) int {
	if row == 0 {
		return 0
	}
	count := 0
	for y := row - 1; y >= 0; y-- {
		if forest[y][col] < tree {
			count++
		} else if forest[y][col] >= tree {
			count++
			break
		}
	}
	return count
}

func visIndexDown(forest [][]int, col int, row int, tree int, height int) int {
	if row == height-1 {
		return 0
	}
	count := 0
	for y := row + 1; y < height; y++ {
		if forest[y][col] < tree {
			count++
		} else if forest[y][col] >= tree {
			count++
			break
		}
	}
	return count
}

func checkVisibility(forest [][]int, col, row, tree, width, height int) int {
	//fmt.Println(col, row, tree, width, height)
	if col == 0 || row == 0 || col == width-1 || row == height-1 {
		return 1
	}
	if checkLeft(forest, col, row, tree) || checkRight(forest, col, row, width, tree) ||
		checkTop(forest, col, row, tree) || checkBottom(forest, col, row, height, tree) {
		return 1
	}
	return 0
}

func checkBottom(forest [][]int, col int, row int, height int, tree int) bool {
	for y := row + 1; y < height; y++ {
		if forest[y][col] >= tree {
			//fmt.Println("ups higher tree than", tree, "at", y, col, "=", forest[y][col])
			return false
		}
	}
	return true
}

func checkTop(forest [][]int, col int, row int, tree int) bool {
	for y := 0; y < row; y++ {
		if forest[y][col] >= tree {
			//fmt.Println("ups higher tree than", tree, "at", y, col, "=", forest[y][col])
			return false
		}
	}
	return true
}

func checkRight(forest [][]int, col int, row int, width int, tree int) bool {
	for x := col + 1; x < width; x++ {
		if forest[row][x] >= tree {
			//fmt.Println("pos higher tree than", tree, "at", x, row, "=", forest[row][x])
			return false
		}
	}
	return true
}

func checkLeft(forest [][]int, col int, row int, tree int) bool {
	for x := 0; x < col; x++ {
		if forest[row][x] >= tree {
			//fmt.Println("pre higher tree than", tree, "at", x, row, "=", forest[row][x])
			return false
		}
	}
	return true
}

func spliceAndParse(line string) []int {
	charas := strings.Split(line, "")
	var row []int
	for _, num := range charas {
		val, _ := strconv.Atoi(num)
		row = append(row, val)
	}
	return row
}
