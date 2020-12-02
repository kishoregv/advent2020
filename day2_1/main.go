package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type pass struct {
	lb    int
	ub    int
	c     rune
	p     string
	count int
	valid bool
}

func main() {
	passes, err := readInput()
	if err != nil {
		panic(err)
	}
	validPasses := 0
	validatePass(passes)
	for _, p := range passes {
		if p.valid {
			fmt.Printf("pass valid %+v, char is %c\n", p, p.c)
			validPasses++
			continue
		}
		fmt.Printf("pass invalid %+v, char is %c\n", p, p.c)
	}
	fmt.Printf("valid passes count %d", validPasses)
}

func readInput() ([]*pass, error) {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return nil, errors.New("failed to read the input file")
	}
	inputStr := string(b)
	numbers := []*pass{}
	for _, in := range strings.Split(inputStr, "\n") {
		p := &pass{}
		parts := strings.Split(in, " ")
		p.p = parts[2]
		p.c = rune(parts[1][0])
		bounds := strings.Split(parts[0], "-")
		lb, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, fmt.Errorf("lb failed %v", err)
		}
		p.lb = lb
		ub, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, fmt.Errorf("ub failed %v", err)
		}
		p.ub = ub

		numbers = append(numbers, p)
	}
	return numbers, nil
}

func validatePass(in []*pass) {
	for _, p := range in {
		for _, c := range p.p {
			if c == p.c {
				p.count++
			}
		}
		if p.count >= p.lb && p.count <= p.ub {
			p.valid = true
		}
	}
}
