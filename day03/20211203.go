package day03

import (
	"fmt"
	"github.com/ChristianGaertner/aoc2021/common"
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
	return "2021 12 03"
}

func SolvePartOne() error {
	numbers, err := common.LoadSValueLines("data/03.txt")
	if err != nil {
		return err
	}

	fmt.Println("Part One")

	var gamma []string
	var epsilon []string

	for i := 0; i < len(numbers[0]); i++ {

		a, b := mostCommonAtPos(numbers, i)

		gamma = append(gamma, a)
		epsilon = append(epsilon, b)
	}

	fmt.Println(toDecimal(gamma) * toDecimal(epsilon))

	return nil
}

func SolvePartTwo() error {
	numbers, err := common.LoadSValueLines("data/03.txt")
	if err != nil {
		return err
	}

	ox := f(numbers, 0, true)
	sc := f(numbers, 0, false)

	fmt.Println(ox.IntBaseN(2) * sc.IntBaseN(2))

	return nil
}

func f(lines []common.SValue, at int, cmmn bool) common.SValue {
	if len(lines) == 1 {
		return lines[0]
	}
	a, b := mostCommonAtPos(lines, at)

	c := b
	if cmmn {
		c = a
	}

	var res []common.SValue
	for _, l := range lines {
		if string(l[at]) == c {
			res = append(res, l)
		}
	}
	return f(res, at+1, cmmn)
}

func mostCommonAtPos(lines []common.SValue, at int) (string, string) {
	var c0, c1 int
	for _, n := range lines {
		if n[at] == '0' {
			c0++
		} else {
			c1++
		}
	}
	if c0 > c1 {
		return "0", "1"
	} else {
		return "1", "0"
	}
}

func toDecimal(bits []string) int {
	g, err := strconv.ParseInt(strings.Join(bits, ""), 2, 32)
	if err != nil {
		panic(err)
	}
	return int(g)
}

func toDecimalS(bits string) int {
	g, err := strconv.ParseInt(bits, 2, 32)
	if err != nil {
		panic(err)
	}
	return int(g)
}
