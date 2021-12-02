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
		switch ins.Str(0) {
		case "forward":
			x += ins.Int(1)
		case "down":
			z += ins.Int(1)
		case "up":
			z -= ins.Int(1)
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
		switch ins.Str(0) {
		case "forward":
			x += ins.Int(1)
			z += aim * ins.Int(1)
		case "down":
			aim += ins.Int(1)
		case "up":
			aim -= ins.Int(1)
		}
	}

	fmt.Println(x * z)

	return nil
}

