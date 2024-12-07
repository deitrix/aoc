package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day7"
)

type Calibration struct {
	Expected int
	Operands []int
}

func parseCalibrations(lines []string) []Calibration {
	calibrations := make([]Calibration, len(lines))
	for i, line := range lines {
		calibrations[i] = parseCalibration(line)
	}
	return calibrations
}

func main() {
	calibrations := parseCalibrations(slices.Collect(aoc.Lines(day7.Input)))
	var maxOperands int
	var maxOperand int
	for _, cal := range calibrations {
		maxOperands = max(maxOperands, len(cal.Operands))
		for _, operand := range cal.Operands {
			maxOperand = max(maxOperand, operand)
		}
	}

	// pre-compute log10 and pow10 results, as it gets called millions of times.
	log10Cache = make([]int, maxOperand+1)
	pow10Cache = make([]int, maxOperands+1)
	for _, cal := range calibrations {
		for _, operand := range cal.Operands[1:] {
			if log10Cache[operand] == 0 {
				log10Cache[operand] = log10(operand)
				pow10Cache[log10Cache[operand]] = pow(10, log10Cache[operand])
			}
		}
	}

	var result int
	for _, cal := range calibrations {
		if hasSolution(cal.Expected, cal.Operands) {
			result += cal.Expected
		}
	}

	fmt.Printf("Result: %d\n", result)
}

// hasSolution returns the first permutation of operators that results in the expected answer.
// It returns -1 if no solution is found.
func hasSolution(expected int, operands []int) bool {
	for operators := range operatorPermutations(len(operands) - 1) {
		if calculate(operands, operators) == expected {
			return true
		}
	}

	return false
}

const operatorCount = 3

func operatorPermutations(size int) iter.Seq[[]int] {
	s := make([]int, size)
	return func(yield func([]int) bool) {
		var recurse func(index int) bool
		recurse = func(index int) bool {
			for x := range operatorCount {
				s[index] = x
				if index < size-1 {
					if !recurse(index + 1) {
						return false
					}
				} else {
					if !yield(s) {
						return false
					}
				}
			}
			return true
		}
		recurse(0)
	}
}

// calculate applies the given permutation of operators to the given operands. It returns the answer
// to the expression.
func calculate(operands []int, operators []int) int {
	answer := operands[0]
	for j, operand := range operands[1:] {
		switch operators[j] {
		case 0:
			answer = answer + operand
		case 1:
			answer = answer * operand
		case 2:
			answer = answer*pow10Cache[log10Cache[operand]] + operand
		}
	}
	return answer
}

func parseCalibration(line string) (cal Calibration) {
	fields := strings.Fields(line)
	cal.Expected = aoc.Int(fields[0][:len(fields[0])-1])
	cal.Operands = make([]int, len(fields[1:]))
	for i, field := range fields[1:] {
		cal.Operands[i] = aoc.Int(field)
	}
	return cal
}

var pow10Cache []int

// pow is faster than math.Pow for small powers.
func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	result := a
	for i := 1; i < b; i++ {
		result *= a
	}
	return result
}

var log10Cache []int

// log10 returns the number of digits in the given number.
func log10(n int) int {
	l := 0
	for n > 0 {
		n /= 10
		l++
	}
	return l
}
