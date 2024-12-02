package main

import (
	"fmt"
	"os"
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
		if safe {
			safeCount++
		}

		fmt.Fprintf(tw, "%s\tsafe = %t\tsafeCount = %d\n", line, safe, safeCount)
	}
}

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
