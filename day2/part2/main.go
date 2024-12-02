package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"

	. "github.com/deitrix/aoc"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	var safeCount int
	for line := range Lines() {
		fields := strings.Fields(line)
		Assert(len(fields) > 1, "Expected at least 2 fields")

		levels := make([]int, len(fields))
		for i, fld := range fields {
			levels[i] = Must1(strconv.Atoi(fld))
		}

		safe := isSafe(levels)
		dampen := -1 // the index of a level to remove to make it safe
		if !safe {
			// If not safe, see if there's a single level we could remove to make it safe.
			for i := 0; i < len(levels); i++ {
				safe = isSafe(slices.Delete(slices.Clone(levels), i, i+1))
				if safe {
					dampen = i
					break
				}
			}
		}
		if safe {
			safeCount++
		}

		if dampen > -1 {
			fmt.Fprintf(tw, "%s\tsafe = %t (remove [%d])\tsafeCount = %d\n", line, safe, dampen, safeCount)
		} else {
			fmt.Fprintf(tw, "%s\tsafe = %t\tsafeCount = %d\n", line, safe, safeCount)
		}
	}
}

func isSafe(levels []int) bool {
	asc := levels[1] > levels[0]
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff > 0 != asc {
			return false // mixture of increasing and decreasing; not safe
		}
		if a := abs(diff); a == 0 || a > 3 {
			return false // too little or too much of a jump; not safe
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
