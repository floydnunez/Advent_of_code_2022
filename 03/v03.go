package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("03/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	var wrongs []string
	for _, line := range lines {
		length := len(strings.TrimSpace(line))
		halflen := length / 2
		fh := line[:halflen]
		sh := line[halflen:]
		both := letterInBoth(fh, sh)
		bothstr := string(both)
		//fmt.Println("both:", bothstr, both)
		wrongs = append(wrongs, bothstr)
	}
	totalP1 := calcTotal(wrongs)
	fmt.Println("Part 1:", totalP1) //8202

	//Part 2
	total := len(lines)
	var badges []string
	for i := 0; i < total/3; i++ {
		l1 := lines[i*3]
		l2 := lines[i*3+1]
		l3 := lines[i*3+2]
		troth := letterInAllThree(l1, l2, l3)
		trothstr := string(troth)
		//println(l1)
		//println(l2)
		//println(l3)
		badges = append(badges, trothstr)
		//println("troth:", trothstr)
	}
	totalP2 := calcTotal(badges)
	fmt.Println("Part 2:", totalP2) //2864
}

func calcTotal(wrongs []string) int {
	totalP1 := 0
	for _, wrong := range wrongs {
		prior := wrong[0] + 1
		if unicode.IsUpper(rune(wrong[0])) {
			prior -= 'A'
			prior += 26
		} else {
			prior -= 'a'
		}
		//fmt.Println(wrong, "prior:", prior)
		totalP1 += int(prior)
	}
	return totalP1
}

func letterInAllThree(l1 string, l2 string, l3 string) int32 {
	for _, c1 := range l1 {
		for _, c2 := range l2 {
			for _, c3 := range l3 {
				if c1 == c2 && c2 == c3 {
					return c1
				}
			}
		}
	}
	return -1
}

func letterInBoth(fh string, sh string) int32 {
	for _, letter := range fh {
		for _, secondLetter := range sh {
			if letter == secondLetter {
				return letter
			}
		}
	}
	return -1
}
