package day01

import (
	"fmt"
	"github.com/ChristianGaertner/aoc2021/common"
)

type Solver struct{}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2021 12 01"
}

func SolvePartOne() error {
	numbers, err := common.LoadMValueLines("data/01.txt")
	if err != nil {
		return err
	}

	var inc int

	fmt.Println("Part One")
	for i, a := range numbers {
		if i > 0 && a.Int(0) > numbers[i-1].Int(0) {
			inc++
		}
	}

	fmt.Println(inc)

	return nil
}

func SolvePartTwo() error {
	numbers, err := common.LoadMValueLines("data/01.txt")
	if err != nil {
		return err
	}

	ws := 3

	var windows []int

	fmt.Println("Part Two")
	for w := 0; w+ws-1 < len(numbers); w++ {
		s := numbers[w].Int(0) + numbers[w+1].Int(0) + numbers[w+2].Int(0)
		windows = append(windows, s)
	}

	var inc int

	for i, a := range windows {
		if i > 0 && a > windows[i-1] {
			inc++
		}
	}

	fmt.Println(inc)

	return nil
}
