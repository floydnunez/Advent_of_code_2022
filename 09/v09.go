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

type pos struct {
	name string
	x, y int
}

func main() {
	file := "real"
	dat, err := os.ReadFile("09/" + file + ".txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	println(lines)
	total := 0
	var matrix [][]string
	var side int
	h, t, matrix := initializeMatrixPos(file, side, matrix)
	printM(matrix, file, h, t)
	maxx := 0
	minx := 99999999
	maxy := 0
	miny := 99999999
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])
		fmt.Println("moving towards", parts[0], "by", amount)
		for steps := 0; steps < amount; steps++ {

			h = updateH(parts[0], 1, h)
			matrix, t = updateT(matrix, h, t, true)
		}
		printM(matrix, file, h, t)
		maxx, maxy, minx, miny = calcMaxes(h, maxx, maxy, minx, miny)
	}
	printM(matrix, file, h, t)
	total = countHash(matrix)
	fmt.Println("h:", h, minx, miny, maxx, maxy)
	fmt.Println("Part 1:", total) //6269
	h, t, matrix = initializeMatrixPos(file, side, matrix)
	printM(matrix, file, h, t)
	t1 := copyPos(h)
	t2 := copyPos(h)
	t3 := copyPos(h)
	t4 := copyPos(h)
	t5 := copyPos(h)
	t6 := copyPos(h)
	t7 := copyPos(h)
	t8 := copyPos(h)
	t9 := copyPos(h)
	fmt.Println(len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])
		fmt.Println("p2: moving towards", parts[0], "by", amount)
		for steps := 0; steps < amount; steps++ {
			h = updateH(parts[0], 1, h)
			matrix, t1 = updateT(matrix, h, t1, false)
			matrix, t2 = updateT(matrix, t1, t2, false)
			matrix, t3 = updateT(matrix, t2, t3, false)
			matrix, t4 = updateT(matrix, t3, t4, false)
			matrix, t5 = updateT(matrix, t4, t5, false)
			matrix, t6 = updateT(matrix, t5, t6, false)
			matrix, t7 = updateT(matrix, t6, t7, false)
			matrix, t8 = updateT(matrix, t7, t8, false)
			matrix, t9 = updateT(matrix, t8, t9, true)
		}
		printM(matrix, file, h, t)
		maxx, maxy, minx, miny = calcMaxes(h, maxx, maxy, minx, miny)
	}
	total = countHash(matrix)
	fmt.Println("Part 2:", total) //2557

}

func copyPos(h pos) pos {
	result := pos{name: h.name, x: h.x, y: h.y}
	return result
}

func initializeMatrixPos(file string, side int, matrix [][]string) (pos, pos, [][]string) {
	var h, t pos
	if file == "real" {
		side = 600
		h = pos{name: "H", x: side / 2, y: side / 2}
		t = pos{name: "T", x: side / 2, y: side / 2}
	} else {
		side = 30
		h = pos{name: "H", x: side / 2, y: side / 2}
		t = pos{name: "T", x: side / 2, y: side / 2}
	}
	matrix = nil
	for y := 0; y < side; y++ {
		matrix = append(matrix, nil)
		for x := 0; x < side; x++ {
			matrix[y] = append(matrix[y], ".")
		}
	}
	return h, t, matrix
}

func countHash(matrix [][]string) int {
	height := len(matrix)
	width := len(matrix[0])
	total := 0
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if matrix[y][x] == "#" {
				total += 1
			}
		}
	}
	return total
}

func updateT(matrix [][]string, h pos, t pos, mark bool) ([][]string, pos) {
	if mark {
		matrix[t.y][t.x] = "#"
	}
	if !tooFarAway(h, t) {
		return matrix, t
	}
	if h.x == t.x { //linear mov
		if h.y > t.y {
			t.y += 1
		} else if h.y < t.y { //maybe they are in the same place! don't trust the plain else
			t.y -= 1
		}
	} else if h.y == t.y { //linear mov
		if h.x > t.x {
			t.x += 1
		} else if h.x < t.x { //maybe they are in the same place! don't trust the plain else
			t.x -= 1
		}
	} else if h.x != t.x && h.y != t.y { //diagonals. again, don't trust
		if h.x > t.x && h.y > t.y { //right up
			t.x += 1
			t.y += 1
		}
		if h.x > t.x && h.y < t.y { //right down
			t.x += 1
			t.y -= 1
		}
		if h.x < t.x && h.y > t.y { //left up
			t.x -= 1
			t.y += 1
		}
		if h.x < t.x && h.y < t.y { //left down
			t.x -= 1
			t.y -= 1
		}
	}
	if mark {
		matrix[t.y][t.x] = "#" //twice cause the pos after moving is needed too
	}
	return matrix, t
}

func tooFarAway(h pos, t pos) bool {
	xdis := h.x - t.x
	ydis := h.y - t.y
	if absInt(xdis) > 1 || absInt(ydis) > 1 {
		return true
	}
	return false
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

//
//func updateT(matrix [][]string, s string, amount int, h pos, t pos) ([][]string, pos) {
//
//}

func calcMaxes(h pos, maxx int, maxy int, minx int, miny int) (int, int, int, int) {
	if h.x < 0 || h.y < 0 {
		fmt.Println("ILLEGAL POSITION")
		panic("aaaa")
	}
	if maxx < h.x {
		maxx = h.x
	}
	if maxy < h.y {
		maxy = h.y
	}
	if minx > h.x {
		minx = h.x
	}
	if miny > h.y {
		miny = h.y
	}
	return maxx, maxy, minx, miny
}

func updateH(dir string, amount int, h pos) pos {
	if dir == "R" {
		h.x += amount
	}
	if dir == "L" {
		h.x -= amount
	}
	if dir == "U" {
		h.y += amount
	}
	if dir == "D" {
		h.y -= amount
	}
	return h
}

func printM(matrix [][]string, file string, h pos, t pos) {
	if file == "real" {
		return
	}
	height := len(matrix)
	width := len(matrix[0])
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			if t.x == x && t.y == y {
				print("T")
			} else if h.x == x && h.y == y {
				print("H")
			} else {
				print(matrix[y][x])
			}
		}
		println("")
	}
	println("")
}
