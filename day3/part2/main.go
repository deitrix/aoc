package main

import (
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"text/tabwriter"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day3"
)

var (
	doRe   = regexp.MustCompile(`do\(\)`)
	dontRe = regexp.MustCompile(`don't\(\)`)
	mulRe  = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
)

func main() {
	// Collect up all the locations of the do, don't, and mul instructions
	doLocs := doRe.FindAllIndex(day3.Input, -1)
	dontLocs := dontRe.FindAllIndex(day3.Input, -1)
	mulLocs := mulRe.FindAllSubmatchIndex(day3.Input, -1)

	// Combine all the locations into a single slice of statements.
	stmts := make([]Stmt, 0, len(doLocs)+len(dontLocs)+len(mulLocs))
	for _, loc := range doLocs {
		stmts = append(stmts, Do(loc[0]))
	}
	for _, loc := range dontLocs {
		stmts = append(stmts, Dont(loc[0]))
	}
	for _, loc := range mulLocs {
		a := aoc.Int(string(day3.Input[loc[2]:loc[3]]))
		b := aoc.Int(string(day3.Input[loc[4]:loc[5]]))
		stmts = append(stmts, Mul(loc[0], a, b))
	}

	// Sort the statements by their location in the input
	slices.SortFunc(stmts, func(a, b Stmt) int {
		return cmp.Compare(a[1], b[1])
	})

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	fmt.Fprintln(tw, "Index\tInput\tState\tResult\tTotal")
	var total int
	enabled := true

	// Execute each statement in order and print the results
	for _, stmt := range stmts {
		result := ""
		switch {
		case stmt[0] == 0:
			enabled = true
		case stmt[0] == 1:
			enabled = false
		case stmt[0] == 2 && enabled:
			result = strconv.Itoa(stmt[2] * stmt[3])
			total += stmt[2] * stmt[3]
		}
		state := "Disabled"
		if enabled {
			state = "Enabled"
		}
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t%d\n", stmt[1], stmt, state, result, total)
	}
}

// Stmt is a single instruction in the input program. It is a 4-tuple where the first element is the
// type of instruction, the second element is the location of the instruction in the input, and the
// third and fourth elements are the arguments to the instruction.
type Stmt [4]int

func (s Stmt) String() string {
	switch s[0] {
	case 0:
		return "do()"
	case 1:
		return "don't()"
	case 2:
		return fmt.Sprintf("mul(%d, %d)", s[2], s[3])
	default:
		panic("unknown statement type")
	}
}

func Do(index int) Stmt        { return Stmt{0, index, 0, 0} }
func Dont(index int) Stmt      { return Stmt{1, index, 0, 0} }
func Mul(index, a, b int) Stmt { return Stmt{2, index, a, b} }
