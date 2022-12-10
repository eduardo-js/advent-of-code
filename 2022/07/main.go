package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/eduardo-js/advent-of-code/2022/helper"
)

func main() {
	exampleInput := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
	s := strings.Split(exampleInput, "\n")
	fmt.Printf("part1 example: %d, should be: %d\n", part1(s, 100000), 95437)
	fmt.Printf("part2 example: %d, should be: %d\n", part2(s, 30000000), 24933642)
	arr := helper.GetInput("07")
	fmt.Printf("part1 input: %d\n", part1(arr, 100000))
	fmt.Printf("part2 input: %d\n", part2(arr, 30000000))
}

type Node struct {
	name    string
	size    int64
	subDirs map[string]*Node
	prevDir *Node
}

func mapDirs(str []string) Node {
	var baseDir Node = Node{"/", 0, make(map[string]*Node), nil}
	currentDir := &baseDir
	for i := 1; i < len(str); i++ {
		v := strings.Split(str[i], " ")
		if v[0] == "dir" {
			currentDir.subDirs[v[1]] = &Node{v[1], 0, make(map[string]*Node), currentDir}
		} else if v[1] == "cd" {
			if v[2] == ".." {
				currentDir = currentDir.prevDir
			} else {
				currentDir = currentDir.subDirs[v[2]]
			}
		} else {
			v, _ := strconv.ParseInt(v[0], 10, 64)
			if v != 0 {
				currentDir.size = currentDir.size + v
			}
		}
	}
	return baseDir
}

func part1(str []string, limit int64) int {
	baseDir := mapDirs(str)
	var p int64 = 0
	return int(aggregateDirSize(&baseDir, limit, &p))
}

func aggregateDirSize(dir *Node, limit int64, total *int64) (size int64) {
	for _, v := range dir.subDirs {
		size := aggregateDirSize(v, limit, total)
		dir.size = dir.size + size
	}
	if dir.size <= limit {
		*total += dir.size
		return dir.size
	}
	return *total
}

func sumDirsSizes(dir *Node, dirSizes *[]int) int64 {
	for _, v := range dir.subDirs {
		size := sumDirsSizes(v, dirSizes)
		dir.size = dir.size + size
	}
	*dirSizes = append(*dirSizes, int(dir.size))
	return dir.size
}

const fileSystemSize = 70000000

func part2(str []string, limit int) int {
	baseDir := mapDirs(str)
	dirSizes := []int{}
	_ = int(sumDirsSizes(&baseDir, &dirSizes))
	min := limit - (fileSystemSize - dirSizes[len(dirSizes)-1])
	sort.Ints(dirSizes)
	for _, v := range dirSizes {
		if v >= min {
			return v
		}
	}
	return 0
}
