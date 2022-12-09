package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type file struct {
	name string
	size int
}

type dir struct {
	file
	files  map[string]file
	dirs   map[string]dir
	parent *dir
}

func newDir(name string) dir {
	result := dir{file: file{name: name, size: 0}}
	result.files = make(map[string]file)
	result.dirs = make(map[string]dir)
	return result
}

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Couldn't convert to int: ", str, err)
		panic("Couldn't convert to int: " + str)
	}
	return val
}

func main() {
	dat, err := os.ReadFile("07/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	root := newDir("/")
	root.parent = &root
	current := &root
	fileRegex := regexp.MustCompile(`\d \w`)
	for _, line := range lines {
		println(line)
		parts := strings.Split(line, " ")
		if line == "$ cd /" || line == "$ ls" {
			continue
		} else if strings.HasPrefix(line, "dir") {
			thisDir := newDir(parts[1])
			thisDir.parent = current
			current.dirs[parts[1]] = thisDir
			fmt.Println("1 this dir", thisDir)
			fmt.Println("1 current dirs", current.dirs)
			if len(thisDir.name) <= 0 {
				fmt.Println("error! empty dir?", line)
				panic("empty name")
			}
		} else if line == "$ cd .." {
			fmt.Println(" 4 cd .. current:", current)
			fmt.Println(" 4 cd .. parent:", current.parent)
			current = current.parent
		} else if strings.HasPrefix(line, "$ cd") {
			fmt.Println(" 2  going to cd to one of:", current.dirs)
			thisDir := current.dirs[parts[2]]
			thisDir.parent = current
			current = &thisDir
			fmt.Println(" 2  cd to this dir:", thisDir)
		} else if fileRegex.MatchString(line) {
			thisFile := file{name: parts[1], size: toInt(parts[0])}
			current.files[thisFile.name] = thisFile
			current.size += thisFile.size
			fmt.Println("  3   current dir:", current.name)
			fmt.Println("  3   file:", thisFile.name)
			fmt.Println("  3   size:", thisFile.size)
			fmt.Println("  3   current.dirs:", current.dirs)
			fmt.Println("  3   current.files:", current.files)
		}
	}
	root.size = cumulativeSize(&root)
	fmt.Println("root:", root)
	smalldirs := findSmallDirs(root)
	part1 := sumtotal(smalldirs)
	fmt.Println("Part 1 smalldirs:", smalldirs)
	fmt.Println("Part 1:", part1)

}

func sumtotal(smalldirs []dir) int {
	total := 0
	for _, val := range smalldirs {
		total += val.size
	}
	return total
}

func findSmallDirs(root dir) []dir {
	var result []dir
	for _, val := range root.dirs {
		fmt.Println("small dirs:", val.name, val.size)
		if len(val.name) <= 0 {
			fmt.Println("woah", val)
		}
		if val.size <= 100000 {
			result = append(result, val)
			result = append(result, findSmallDirs(val)...)
		}
	}
	return result
}

func cumulativeSize(current *dir) int {
	total := 0
	for _, v := range current.files {
		total += v.size
	}
	for _, v := range current.dirs {
		total += cumulativeSize(&v)
	}
	current.size = total
	myself := current.parent.dirs[current.name]
	myself.size = total
	current.parent.dirs[current.name] = myself
	return total
}
