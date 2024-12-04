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
	width := len(lines[0])
	height := len(lines)

	// First, find all occurrences of the letter X
	var Xes [][2]int
	for y, line := range lines {
		for x, ch := range line {
			if ch == 'X' {
				Xes = append(Xes, [2]int{x, y})
			}
		}
	}

	// Then, for each occurrence of X, fan out in all 8 direction, searching for "XMAS", and
	// keeping track of the coordinates so that we can pretty print them in the next step
	indices := make(map[[2]int]struct{})
	var count int
	for _, pos := range Xes {
		for _, d := range dir {
			x, y := pos[0], pos[1]
			coords := [][2]int{{x, y}}
			for i := 1; i < 4; i++ {
				x += d[0]
				y += d[1]
				coords = append(coords, [2]int{x, y})
				if x < 0 || x >= width || y < 0 || y >= height {
					break
				}
				if lines[y][x] != "XMAS"[i] {
					break
				}
				if i == 3 {
					count++
					for _, c := range coords {
						indices[c] = struct{}{}
					}
				}
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

	fmt.Printf("\nFound %s occurrences of %s\n", color.GreenString("%d", count), color.RedString("XMAS"))
}

// dir contains a list of x and y offsets for each of the 8 directions
var dir = [][2]int{
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}
