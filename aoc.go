package aoc

import (
	"bufio"
	"iter"
	"os"
	"strconv"
	"strings"
)

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func Lines() iter.Seq[string] {
	return func(yield func(string) bool) {
		f, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(f)
		defer f.Close()
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}

func Ints(line string) []int {
	fields := strings.Fields(line)
	ints := make([]int, len(fields))
	for i, fld := range fields {
		ints[i] = Must1(strconv.Atoi(fld))
	}
	return ints
}

func Must1[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func Abs[T int | uint](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
