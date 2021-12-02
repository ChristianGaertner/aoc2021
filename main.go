package main

import (
	"github.com/ChristianGaertner/aoc2021/common"
	"github.com/ChristianGaertner/aoc2021/day02"
)

func main() {
	s := common.WithTiming(day02.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
