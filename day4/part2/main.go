package main

import (
	"fmt"
	"slices"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day4"
	"github.com/fatih/color"
)

func main() {
	lines := slices.Collect(aoc.Lines(day4.Input))

	// Find all occurrences of the letter A, then check the surrounding 4 squares for the
	// configurations of MAS that we're looking for. Add the coordinates of the MAS squares
	// to a set so that we can pretty print them in the next step.
	indices := make(map[[2]int]struct{})
	var count int
	for y := 1; y < len(lines)-1; y++ {
		line := lines[y]
		for x := 1; x < len(line)-1; x++ {
			if line[x] != 'A' {
				continue
			}
			config := [4]byte{
				lines[y-1][x-1],
				lines[y-1][x+1],
				lines[y+1][x-1],
				lines[y+1][x+1],
			}
			if slices.Contains(configs, config) {
				count++
				indices[[2]int{x + 0, y + 0}] = struct{}{}
				indices[[2]int{x - 1, y - 1}] = struct{}{}
				indices[[2]int{x + 1, y - 1}] = struct{}{}
				indices[[2]int{x - 1, y + 1}] = struct{}{}
				indices[[2]int{x + 1, y + 1}] = struct{}{}
			}
		}
	}

	// Finally, print out the board, highlighting the XMAS coordinates
	for y, line := range lines {
		for x, ch := range line {
			if _, ok := indices[[2]int{x, y}]; ok {
				fmt.Print(color.RedString(string(ch)))
			} else {
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}

	fmt.Printf("\nFound %s occurrences of X-%s\n", color.GreenString("%d", count), color.RedString("MAS"))
}

// configs is a list of all possible configurations of overlapping MAS squares
var configs = [][4]byte{
	{'M', 'S', 'M', 'S'},
	{'S', 'S', 'M', 'M'},
	{'S', 'M', 'S', 'M'},
	{'M', 'M', 'S', 'S'},
}
