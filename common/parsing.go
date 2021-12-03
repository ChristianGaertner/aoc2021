package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MValue []SValue

type SValue string

func (s SValue) Str() string {
	return string(s)
}

func (s SValue) Int() int {
	return s.IntBaseN(10)
}

func (s SValue) Int64() int64 {
	n, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func (s SValue) IntBaseN(base int) int {
	n, err := strconv.ParseInt(string(s), base, 32)
	if err != nil {
		panic(err)
	}
	return int(n)
}

func (s SValue) Float() float64 {
	n, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		panic(err)
	}
	return n
}

func (m MValue) Strings() []string {
	var r []string
	for _, s := range m {
		r = append(r, string(s))
	}
	return r
}

func (m MValue) Ints() []int {
	var res []int
	for _, v := range m {
		res = append(res, v.Int())
	}
	return res
}

func (m MValue) Int64s() []int64 {
	var res []int64
	for _, v := range m {
		res = append(res, v.Int64())
	}
	return res
}

func (m MValue) Floats() []float64 {
	var res []float64
	for _, v := range m {
		res = append(res, v.Float())
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

		m := make(MValue, len(parts))
		for i, p := range parts {
			m[i] = SValue(p)
		}

		res = append(res, m)
	}
	return res, nil
}

func LoadSValueLines(path string) ([]SValue, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var res []SValue

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, SValue(scanner.Text()))
	}
	return res, nil
}
