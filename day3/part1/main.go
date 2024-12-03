package main

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day3"
)

var mulRe = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	fmt.Fprintln(tw, "Input\tResult\tTotal")
	var total int
	for _, match := range mulRe.FindAllStringSubmatch(string(day3.Input), -1) {
		result := aoc.Int(match[1]) * aoc.Int(match[2])
		total += result
		fmt.Fprintf(tw, "%s\t%d\t%d\n", match[0], result, total)
	}
}
