package main

import (
	"bytes"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day6"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
)

// OK so my initial thoughts on a solution are to have four separate grids keeping track of the
// patrolled positions for the guard - one for each direction of travel. Then, when we're simulating
// we can check if the next position intersects with the grid for the next direction of travel. If
// it does, placing an obstruction ahead of the guard will cause it to turn onto an already
// established path, putting it in a loop.
//
// Scratch that. Instead, we're just going to simulate the guard's movements as normal, but for each
// movement we'll add an obstruction directly in front of the guard, and start a sub-simulation with
// the new obstruction in mind. If the guard ever intersects with a patrolled position in the same
// direction, we know we've found a loop. This should be more exhaustive than the previous approach.
func main() {
	lines := slices.Collect(aoc.Lines(day6.Input))
	size := len(lines)
	grid := make([]byte, 0, size*size)
	for _, line := range lines {
		grid = append(grid, line...)
	}
	// Find the starting position and direction of the guard.
	var dir int
	var x, y int
FindGuard:
	for i, ch := range grid {
		var ok bool
		if dir, ok = directions[ch]; ok {
			x, y = i%size, i/size
			// The guard doesn't need to exist on the board any more, so to simplify the
			// move logic, we'll replace it with an empty space.
			grid[i] = '.'
			break FindGuard
		}
	}

	// patrolled is a set of positions that the guard has patrolled. The key is a 2-element array
	// containing the index of the position in the grid, and the direction the guard was facing when
	// it patrolled that position.
	patrolled := make(map[[2]int]bool)

	var loopCount int

	uw := uilive.New()
	render := func(grid []byte, size, x, y, dir int, patrolled map[[2]int]bool) {
		buf := new(bytes.Buffer)
		for i, ch := range grid {
			switch {
			case ch == 'O':
				fmt.Fprint(buf, color.CyanString(string(ch)))
			case x == i%size && y == i/size:
				fmt.Fprint(buf, color.WhiteString(string(directionChar[dir])))
			case patrolled[[2]int{i, 0}]:
				fmt.Fprint(buf, directionColor[0]("X"))
			case patrolled[[2]int{i, 1}]:
				fmt.Fprint(buf, directionColor[1]("X"))
			case patrolled[[2]int{i, 2}]:
				fmt.Fprint(buf, directionColor[2]("X"))
			case patrolled[[2]int{i, 3}]:
				fmt.Fprint(buf, directionColor[3]("X"))
			default:
				fmt.Fprint(buf, color.WhiteString(string(ch)))
			}
			if i%size == size-1 {
				fmt.Fprintln(buf)
			}
		}
		fmt.Fprintln(buf)
		fmt.Fprint(uw, buf.String())
		uw.Flush()
	}

	// loopCellBlacklist is a set of grid positions that cannot be considered for placing an
	// obstruction, because either they are the guard's starting position, or they are on the
	// guard's historical path. The obstruction can only be placed at the very start of the
	// simulation, not during. This is to prevent an obstruction from being placed on a position
	// that the guard has already patrolled. If the obstruction was placed from the very start of
	// the simulation, there's no way the guard could have patrolled that position.
	loopCellBlacklist := make(map[int]bool)
	loopCellBlacklist[y*size+x] = true // starting position

	for {
		i := y*size + x
		//render(grid, size, x, y, dir, patrolled)
		//time.Sleep(renderSpeed)

		// Mark the current position as patrolled.
		patrolled[[2]int{i, dir}] = true
		// Also, add it to the loop cell blacklist, since we can no longer place an obstruction
		// here.
		loopCellBlacklist[i] = true

		// Get the coordinates of the next position the guard wants to move to.
		dirVec := ahead[dir]
		nx := x + dirVec[0]
		ny := y + dirVec[1]
		ni := ny*size + nx

		// If the next position is out of bounds, we're done. The guard has left the lab.
		if ny < 0 || ny >= size || nx < 0 || nx >= size {
			break
		}

		// If the next position is an obstruction, turn right.
		if grid[ni] == '#' {
			dir = (dir + 1) % 4
			continue
		}

		if !loopCellBlacklist[ni] {
			// Place an obstruction in front of the guard and simulate the guard's movements to see if
			// it will enter a loop.
			ng := slices.Clone(grid)
			ng[ni] = '#'
			if isCycle(ng, size, x, y, dir, maps.Clone(patrolled), render) {
				loopCount++
				loopCellBlacklist[ni] = true
				grid[ni] = 'O'
			}
		}

		// Move the guard to the next position.
		x, y = nx, ny
	}

	fmt.Printf("Loop count: %d\n", loopCount)
	fmt.Printf("Patrolled %d unique positions\n", len(patrolled))
}

type renderFunc func(grid []byte, size, x, y, dir int, patrolled map[[2]int]bool)

const renderSpeed = time.Second / 100

// isCycle continues the simulation, reporting whether the guard will enter a cycle.
func isCycle(grid []byte, size, x, y, dir int, patrolled map[[2]int]bool, render renderFunc) bool {
	for n := 0; ; n++ {
		i := y*size + x
		//render(grid, size, x, y, dir, patrolled)
		//time.Sleep(renderSpeed)

		// If the guard has patrolled this position in the same direction, we've found a loop. Skip
		// this check for the first iteration, since it will always resolve to true.
		if n > 0 && patrolled[[2]int{i, dir}] {
			return true
		}

		// Mark the current position as patrolled.
		patrolled[[2]int{i, dir}] = true

		// Get the coordinates of the next position the guard wants to move to.
		dirVec := ahead[dir]
		nx := x + dirVec[0]
		ny := y + dirVec[1]
		ni := ny*size + nx

		// If the next position is out of bounds, we're done. The guard has left the lab.
		if ny < 0 || ny >= size || nx < 0 || nx >= size {
			return false
		}

		// If the next position is an obstruction, turn right.
		if grid[ni] == '#' {
			dir = (dir + 1) % 4
			continue
		}

		// Move the guard to the next position.
		x, y = nx, ny
	}
}

// directions maps guard direction characters to their respective integer values.
var directions = map[byte]int{
	'^': 0,
	'>': 1,
	'v': 2,
	'<': 3,
}

var directionChar = "^>v<"

var directionColor = [...]func(string, ...interface{}) string{
	color.BlueString,
	color.RedString,
	color.GreenString,
	color.YellowString,
}

// ahead contains a list of slice index vectors that represent the different directions the guard
// can move. These are in the order of up, right, down, left. Because lines are read from top to
// bottom, the up/down vectors are inverted.
var ahead = [...][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}
