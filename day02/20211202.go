package day02

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
	return "2021 12 02"
}

func SolvePartOne() error {
	instructions, err := common.LoadMValueLines("data/02.txt")
	if err != nil {
		return err
	}

	x, z := 0, 0

	fmt.Println("Part One")
	for _, ins := range instructions {
		switch ins[0].Str() {
		case "forward":
			x += ins[1].Int()
		case "down":
			z += ins[1].Int()
		case "up":
			z -= ins[1].Int()
		}
	}

	fmt.Println(x * z)

	return nil
}

func SolvePartTwo() error {
	instructions, err := common.LoadMValueLines("data/02.txt")
	if err != nil {
		return err
	}

	x, z := 0, 0

	aim := 0

	fmt.Println("Part Two")
	for _, ins := range instructions {
		switch ins[0].Str() {
		case "forward":
			x += ins[1].Int()
			z += aim * ins[1].Int()
		case "down":
			aim += ins[1].Int()
		case "up":
			aim -= ins[1].Int()
		}
	}

	fmt.Println(x * z)

	return nil
}
