package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day7"
	"github.com/fatih/color"
)

const showWork = false

func main() {
	var result int
	for line := range aoc.Lines(day7.Input) {
		expected, operands := parseLine(line)
		solutions := computeSolutions(expected, operands)
		if len(solutions) > 0 {
			result += expected
		}
		if showWork {
			header := fmt.Sprintf("Expected: %d, Operands: %s", expected, aoc.Join(operands, ", "))
			if len(solutions) > 0 {
				header = fmt.Sprintf("%s, Solutions: %d", header, len(solutions))
			}
			sep := strings.Repeat("=", len(header))
			if len(solutions) > 0 {
				color.Green(header)
				fmt.Println(sep)
			} else {
				fmt.Println(header)
				fmt.Println(sep)
			}
			for _, perm := range solutions {
				color.Green(permString(operands, perm))
			}
			for i := 0; i < permCount(len(operands)); i++ {
				if slices.Contains(solutions, i) {
					continue
				}
				fmt.Println(permString(operands, i))
			}
			fmt.Println()
		}
	}
	fmt.Printf("Result: %d\n", result)
}

// computeSolutions returns all the possible solutions to the given expression. It returns a slice
// of integers, each representing a possible solution.
func computeSolutions(expected int, operands []int) []int {
	var solutions []int
	for i := 0; i < permCount(len(operands)); i++ {
		if calculate(operands, i) == expected {
			solutions = append(solutions, i)
		}
	}
	return solutions
}

type operation func(a, b int) int

// ops holds the operations that can be applied to the operands.
var ops = [...]operation{
	0: func(a, b int) int { return a + b },                              // +
	1: func(a, b int) int { return a * b },                              // *
	2: func(a, b int) int { return aoc.Int(fmt.Sprintf("%d%d", a, b)) }, // || (concat)
}

// opStrings holds the string representations of the operations.
var opStrings = [...]string{
	0: "+",
	1: "*",
	2: "||",
}

// calculate applies the given permutation of operators to the given operands. It returns the answer
// to the expression.
func calculate(operands []int, opPerm int) int {
	answer := operands[0]
	for j := 0; j < len(operands)-1; j++ {
		op := opForPerm(j, opPerm)
		answer = ops[op](answer, operands[j+1])
	}
	return answer
}

// permString returns a string representation of the given permutation of operators. It includes
// the operands, operators and the answer: "3 * 3 / 3 = 3".
func permString(operands []int, opPerm int) string {
	buf := new(strings.Builder)
	answer := operands[0]
	fmt.Fprintf(buf, "%d ", operands[0])
	for j := 0; j < len(operands)-1; j++ {
		op := opForPerm(j, opPerm)
		answer = ops[op](answer, operands[j+1])
		fmt.Fprintf(buf, "%s %d ", opStrings[op], operands[j+1])
	}
	fmt.Fprintf(buf, "= %d", answer)
	return buf.String()
}

// permCount returns the number of permutations of the given operators between the operands. Given
// an input of operands: 3, operators: 2, where the two operators correspond to * and /, the
// possible permutations are: (3 * 3 / 3), (3 / 3 * 3), (3 * 3 * 3), (3 / 3 / 3). So, numPerms(3, 2)
// would return 4.
//
// This is simply: operators^(operands-1).
func permCount(operands int) int {
	return int(math.Pow(float64(len(ops)), float64(operands-1)))
}

// opForPerm returns the operation to use at the given index of the permutation.
//
// This is simply: op = (perm / (operators^i)) % operators
func opForPerm(i, perm int) int {
	return (perm / int(math.Pow(float64(len(ops)), float64(i)))) % len(ops)
}

func parseLine(line string) (answer int, operands []int) {
	fields := strings.Fields(line)
	answer = aoc.Int(fields[0][:len(fields[0])-1])
	operands = make([]int, len(fields[1:]))
	for i, field := range fields[1:] {
		operands[i] = aoc.Int(field)
	}
	return answer, operands
}
