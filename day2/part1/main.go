package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/deitrix/aoc"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	var safeCount int
	for line := range aoc.Lines() {
		levels := aoc.Ints(line)
		aoc.Assert(len(levels) > 1, "Expected at least 2 levels")

		safe := isSafe(levels)
		if safe {
			safeCount++
		}

		fmt.Fprintf(tw, "%s\tsafe = %t\tsafeCount = %d\n", line, safe, safeCount)
	}
}

// isSafe reports whether the given reactor levels are safe. In order for the levels to be safe, they must:
// - either be all increasing or all decreasing
// - have a difference of 1, 2, or 3 between each level
func isSafe(levels []int) bool {
	asc := levels[1] > levels[0]
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff > 0 != asc {
			return false // mixture of increasing and decreasing; not safe
		}
		if a := aoc.Abs(diff); a == 0 || a > 3 {
			return false // too little or too much of a jump; not safe
		}
	}
	return true
}
