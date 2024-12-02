package main

import (
	"fmt"
	"os"
	"slices"
	"text/tabwriter"

	. "github.com/deitrix/aoc"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	var safeCount int
	for line := range Lines() {
		levels := Ints(line)
		Assert(len(levels) > 1, "Expected at least 2 levels")

		// First, check if the levels are safe as-is.
		if isSafe(levels) {
			safeCount++
			fmt.Fprintf(tw, "%s\tsafe = true\tsafeCount = %d\n", line, safeCount)
			continue
		}

		// If the levels are not safe, see if we can remove a single level to make them safe.
		if index := dampen(levels); index > -1 {
			safeCount++
			fmt.Fprintf(tw, "%s\tsafe = true (remove [%d])\tsafeCount = %d\n", line, index, safeCount)
			continue
		}

		// The levels are not safe, and we cannot make them safe by removing a single level.
		fmt.Fprintf(tw, "%s\tsafe = false\tsafeCount = %d\n", line, safeCount)
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
		if a := Abs(diff); a == 0 || a > 3 {
			return false // too little or too much of a jump; not safe
		}
	}
	return true
}

// dampen attempts to remove a single level from the list to make it safe. If successful, it returns
// the index of the level to remove, otherwise it returns -1.
func dampen(levels []int) int {
	for i := 0; i < len(levels); i++ {
		if isSafe(slices.Delete(slices.Clone(levels), i, i+1)) {
			return i
		}
	}
	return -1
}
