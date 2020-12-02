package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers, err := readInput()
	if err != nil {
		panic(err)
	}
	sort.Ints(numbers)
	fmt.Println(len(numbers))
	for _, in := range numbers {
		pair, found := binarySearch(in, numbers)
		if found {
			fmt.Printf("found the pair, %d + %d = %d\n %d * %d = %d \n", in, pair, in+pair, in, pair, in*pair)
			return
		}
	}
	fmt.Println("nothing found")
}

func readInput() ([]int, error) {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, errors.New("failed to read the input file")
	}
	inputStr := string(b)
	numbers := []int{}
	for _, in := range strings.Split(inputStr, "\n") {
		num, err := strconv.Atoi(in)
		if err != nil {
			return nil, errors.New("failed to convert string to number")
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func binarySearch(in int, numbers []int) (int, bool) {
	size := len(numbers)
	if size == 0 {
		return 0, false
	}
	if size == 1 {
		if in+numbers[0] == 2020 {
			return numbers[0], true
		}
		return 0, false
	}
	if size == 2 {
		if in+numbers[0] == 2020 {
			return numbers[0], true
		}
		if in+numbers[1] == 2020 {
			return numbers[1], true
		}
		return 0, false
	}
	mid := numbers[size/2]
	sum := mid + in
	if sum > 2020 {
		return binarySearch(in, numbers[0:size/2-1])
	}
	if sum < 2020 {
		return binarySearch(in, numbers[size/2+1:])
	}
	return mid, true
}
