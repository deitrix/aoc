package main

import (
	"bytes"
	"fmt"
	"slices"
	"time"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day6"
	"github.com/gosuri/uilive"
)

func main() {
	lines := slices.Collect(aoc.Lines(day6.Input))
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	// Find the starting position and direction of the guard.
	var dir int
	var x, y int
FindGuard:
	for i, row := range grid {
		for j, ch := range row {
			var ok bool
			if dir, ok = directions[ch]; ok {
				x, y = j, i
				// The guard doesn't need to exist on the board any more, so to simplify the
				// move logic, we'll replace it with an empty space.
				grid[i][j] = '.'
				break FindGuard
			}
		}
	}

	patrolled := make(map[[2]int]bool)

	uw := uilive.New()
	render := func() {
		buf := new(bytes.Buffer)
		for i, row := range grid {
			for j, ch := range row {
				switch {
				case x == j && y == i:
					fmt.Fprint(buf, string(directionChar[dir]))
				case patrolled[[2]int{j, i}]:
					fmt.Fprint(buf, "X")
				default:
					fmt.Fprint(buf, string(ch))
				}
			}
			fmt.Fprintln(buf)
		}
		fmt.Fprintln(buf)
		fmt.Fprint(uw, buf.String())
		uw.Flush()
	}

	for {
		render()
		time.Sleep(time.Second / 100)

		// Mark the current position as patrolled.
		patrolled[[2]int{x, y}] = true

		// Check whether the guard can move in the current direction.
		dirVec := ahead[dir]
		nx := x + dirVec[0]
		ny := y + dirVec[1]

		// If the next position is out of bounds, we're done.
		if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
			break
		}

		// If the next position is an obstruction, turn right.
		if grid[ny][nx] == '#' {
			dir = (dir + 1) % 4
			continue
		}

		// Move the guard to the next position.
		x, y = nx, ny
	}

	fmt.Printf("Patrolled %d unique positions\n", len(patrolled))
}

// directions maps guard direction characters to their respective integer values.
var directions = map[byte]int{
	'^': 0,
	'>': 1,
	'v': 2,
	'<': 3,
}

var directionChar = "^>v<"

// ahead contains a list of slice index vectors that represent the different directions the guard
// can move. These are in the order of up, right, down, left. Because lines are read from top to
// bottom, the up/down vectors are inverted.
var ahead = [...][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}
