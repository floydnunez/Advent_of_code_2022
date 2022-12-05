package main

import (
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
	file := "real"
	dat, err := os.ReadFile("05/" + file + ".txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	var stacks []([]string)
	stacks = initStacks(file, stacks)
	printStacks(stacks)
	for _, line := range lines {
		code := strings.Split(line, " ")
		stacks = moveXfromYtoZ(stacks, code[1], code[3], code[5])
	}
	endconfig := printStacks(stacks)
	printConfig("Part 1:", endconfig) //VPCDMSLWJ
	println("\n")
	//Part 2: reinitialize
	stacks = initStacks(file, stacks)
	printStacks(stacks)
	for _, line := range lines {
		code := strings.Split(line, " ")
		stacks = moveXfromYtoZPart2(stacks, code[1], code[3], code[5])
	}
	endconfig = printStacks(stacks)
	printConfig("Part 2:", endconfig) //TPWCGNCCG

}

func initStacks(file string, stacks [][]string) [][]string {
	if file == "example" {
		stacks = [][]string{
			makeArray("ZN"),
			makeArray("MCD"),
			makeArray("P"),
		}
	} else {
		stacks = [][]string{
			makeArray("NRGP"),
			makeArray("JTBLFGDC"),
			makeArray("MSV"),
			makeArray("LSRCZP"),
			makeArray("PSLVCWDQ"),
			makeArray("CTNWDMS"),
			makeArray("HDGWP"),
			makeArray("ZLPHSCMV"),
			makeArray("RPFLWGZ"),
		}
	}
	return stacks
}

func moveXfromYtoZ(stacks [][]string, s1 string, s2 string, s3 string) [][]string {
	x, _ := strconv.Atoi(s1)
	y, _ := strconv.Atoi(s2)
	z, _ := strconv.Atoi(s3)
	y -= 1
	z -= 1
	from := stacks[y]
	to := stacks[z]
	for i := 0; i < x; i++ {
		popped := from[len(from)-1]
		from = from[:len(from)-1]
		to = append(to, popped)
	}
	stacks[y] = from
	stacks[z] = to
	return stacks
}

func moveXfromYtoZPart2(stacks [][]string, s1 string, s2 string, s3 string) [][]string {
	x, _ := strconv.Atoi(s1)
	y, _ := strconv.Atoi(s2)
	z, _ := strconv.Atoi(s3)
	y -= 1
	z -= 1
	from := stacks[y]
	to := stacks[z]
	popped := from[len(from)-x:]
	from = from[:len(from)-x]
	for _, chara := range popped {
		to = append(to, chara)
	}
	stacks[y] = from
	stacks[z] = to
	return stacks
}

func printConfig(s string, line []string) {
	print(s, " ")
	for _, chara := range line {
		print(chara)
	}
	println("")
}

func printStacks(stacks [][]string) []string {
	var finals []string
	for number, line := range stacks {
		finals = append(finals, "")
		for _, chara := range line {
			print(chara)
			finals[number] = chara
		}
		println("")
	}
	return finals
}

func makeArray(s string) []string {
	return strings.Split(strings.ToUpper(s), "")
}
