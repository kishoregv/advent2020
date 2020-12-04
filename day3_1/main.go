package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	treeMap, err := readInput()
	if err != nil {
		panic(err)
	}
	t1 := countTrees(treeMap, 1, 1)
	t2 := countTrees(treeMap, 3, 1)
	t3 := countTrees(treeMap, 5, 1)
	t4 := countTrees(treeMap, 7, 1)
	t5 := countTrees(treeMap, 1, 2)
	fmt.Printf("number of trees %d", t1*t2*t3*t4*t5)
}

func readInput() ([]string, error) {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, errors.New("failed to read the input file")
	}
	inputStr := string(b)
	return strings.Split(inputStr, "\n"), nil
}

func countTrees(treeMap []string, xslope, yslope int) int {
	cols := len(treeMap[0])
	rows := len(treeMap)
	x := 0
	y := 0
	trees := 0
	for y < rows {
		x = (x + xslope) % cols
		y = y + yslope
		if y >= rows {
			break
		}
		if treeMap[y][x] == '#' {
			trees++
		}
	}
	return trees
}
