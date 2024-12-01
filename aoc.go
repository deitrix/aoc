package aoc

import (
	"bufio"
	"iter"
	"os"
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

func Must1[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}
