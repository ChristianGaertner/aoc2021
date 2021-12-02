package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	instructions, err := _read()
	if err != nil {
		return err
	}


	x, z := 0, 0


	fmt.Println("Part One")
	for _, ins := range instructions {
		switch ins.Dir {
		case "forward":
			x += ins.Amount
		case "down":
			z += ins.Amount
		case "up":
			z -= ins.Amount
		}
	}

	fmt.Println(x * z)

	return nil
}

func SolvePartTwo() error {
	instructions, err := _read()
	if err != nil {
		return err
	}


	x, z := 0, 0

	aim := 0


	fmt.Println("Part Two")
	for _, ins := range instructions {
		switch ins.Dir {
		case "forward":
			x += ins.Amount
			z += aim * ins.Amount
		case "down":
			aim += ins.Amount
		case "up":
			aim -= ins.Amount
		}
	}

	fmt.Println(x * z)

	return nil
}

type Instr struct {
	Dir string
	Amount int
}

func _read() ([]Instr, error) {
	file, err := os.Open("data/02.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var i []Instr

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		parts := strings.Split(raw, " ")



		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		i = append(i, Instr{
			Dir:    parts[0],
			Amount: n,
		})
	}
	return i, nil
}
