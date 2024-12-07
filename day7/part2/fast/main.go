package main

import (
	"fmt"
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
	//f, err := os.Create("cpu.prof")
	//if err != nil {
	//	log.Fatal("could not create CPU profile: ", err)
	//}
	//defer f.Close() // error handling omitted for example
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	log.Fatal("could not start CPU profile: ", err)
	//}
	//defer pprof.StopCPUProfile()

	calibrations := parseCalibrations(slices.Collect(aoc.Lines(day7.Input)))
	var maxOperands int
	var maxOperand int
	var log10Operands int
	for _, cal := range calibrations {
		if len(cal.Operands) > maxOperands {
			maxOperands = len(cal.Operands)
		}
		log10Operands += len(cal.Operands) - 1
		for _, operand := range cal.Operands {
			maxOperand = max(maxOperand, operand)
		}
	}

	// pre-compute opForPerm results, as it gets called millions of times.
	opForPermCache = make([]int, permCount(maxOperands)*(maxOperands-1))
	opForPermWidth = maxOperands - 1
	for i := 0; i < permCount(maxOperands); i++ {
		for j := 0; j < maxOperands-1; j++ {
			opForPermCache[i*opForPermWidth+j] = opForPerm(j, i)
		}
	}

	// pre-compute log10 results, as it gets called millions of times.
	log10Cache = make([]int, log10Operands)
	for _, cal := range calibrations {
		for _, operand := range cal.Operands[1:] {
			log10Cache[operand] = log10(operand)
		}
	}

	results := make(chan int)
	for _, cal := range calibrations {
		go func() {
			solution := firstSolution(cal.Expected, cal.Operands)
			if solution > -1 {
				results <- cal.Expected
			} else {
				results <- 0
			}
		}()
	}

	var result int
	for range len(calibrations) {
		result += <-results
	}

	fmt.Printf("Result: %d\n", result)
}

// firstSolution returns the first permutation of operators that results in the expected answer.
// It returns -1 if no solution is found.
func firstSolution(expected int, operands []int) int {
	for i := 0; i < permCount(len(operands)); i++ {
		if calculate(operands, i) == expected {
			return i
		}
	}

	return -1
}

const ops = 3

// calculate applies the given permutation of operators to the given operands. It returns the answer
// to the expression.
func calculate(operands []int, opPerm int) int {
	answer := operands[0]
	for j := 0; j < len(operands)-1; j++ {
		op := opForPermCache[opPerm*opForPermWidth+j]
		switch op {
		case 0:
			answer = answer + operands[j+1]
		case 1:
			answer = answer * operands[j+1]
		case 2:
			answer = answer*pow(10, log10Cache[operands[j+1]]) + operands[j+1]
		}
	}
	return answer
}

// permCount returns the number of permutations of the given operators between the operands. Given
// an input of operands: 3, operators: 2, where the two operators correspond to * and /, the
// possible permutations are: (3 * 3 / 3), (3 / 3 * 3), (3 * 3 * 3), (3 / 3 / 3). So, numPerms(3, 2)
// would return 4.
//
// This is simply: operators^(operands-1).
func permCount(operands int) int {
	return pow(ops, operands-1)
}

var opForPermCache []int
var opForPermWidth int

// opForPerm returns the operation to use at the given index of the permutation.
//
// This is simply: op = (perm / (operators^i)) % operators
func opForPerm(i, perm int) int {
	return perm / pow(ops, i) % ops
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
