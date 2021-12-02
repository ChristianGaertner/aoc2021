package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MValue []string

func (m MValue) Str(at int) string {
	return m[at]
}

func (m MValue) Int(at int) int {
	n, err := strconv.Atoi(m[at])
	if err != nil {
		panic(err)
	}
	return n
}

func (m MValue) Float(at int) float64 {
	n, err := strconv.ParseFloat(m[at], 64)
	if err != nil {
		panic(err)
	}
	return n
}

func (m MValue) Strings() []string {
	return m
}

func (m MValue) Ints() []int {
	var res []int
	for _, v := range m {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		res = append(res, n)
	}
	return res
}

func (m MValue) Int64s() []int64 {
	var res []int64
	for _, v := range m {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}

		res = append(res, n)
	}
	return res
}

func (m MValue) Floats() []float64 {
	var res []float64
	for _, v := range m {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}

		res = append(res, n)
	}
	return res
}

func LoadMValueLines(path string) ([]MValue, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res []MValue

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		parts := strings.Split(raw, " ")
		res = append(res, parts)
	}
	return res, nil
}
