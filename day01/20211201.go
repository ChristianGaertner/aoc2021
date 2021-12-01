package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	numbers, err := _read()
	if err != nil {
		return err
	}


	var inc int

	fmt.Println("Part One")
	for i, a := range numbers {
		if i > 0 && a > numbers[i - 1] {
			inc++
		}
	}

	fmt.Println(inc)

	return nil
}

func SolvePartTwo() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	ws := 3

	var windows []int

	fmt.Println("Part Two")
	for w := 0; w + ws - 1 < len(numbers); w++ {
		s := numbers[w] + numbers[w+1] + numbers[w+2]
		windows = append(windows, s)
	}

	var inc int

	for i, a := range windows {
		if i > 0 && a > windows[i - 1] {
			inc++
		}
	}

	fmt.Println(inc)

	return nil
}

func _read() ([]int, error) {
	file, err := os.Open("data/01.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		n, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}
