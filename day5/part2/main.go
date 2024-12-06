package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day5"
	"github.com/fatih/color"
)

func main() {
	lines := slices.Collect(aoc.Lines(day5.Input))
	// Find the point where the rules end and the updates begin.
	updatesIdx := slices.Index(lines, "")
	rules := ParseRules(lines[:updatesIdx])
	var total int
	for _, update := range ParseUpdates(lines[updatesIdx+1:]) {
		// Map the sequence values to their indices.
		valid := IsUpdateValid(update, rules)
		if valid {
			fmt.Println(color.WhiteString("%v: Valid", update))
		} else {
			fmt.Print(color.RedString("%v: Invalid => ", update))
			ReorderUpdate(update, rules)
			middle := len(update) / 2
			fmt.Println(color.GreenString("%v: Valid(%d)", update, update[middle]))
			total += update[middle]
		}
	}
	fmt.Printf("Total: %d\n", total)
}

// ReorderUpdate reorders the given update according to the given rules.
func ReorderUpdate(update Update, rules map[[2]int]bool) {
	// For each number in the update, see if any of the other numbers in the update should come
	// before it. If so, swap them. Repeating this for the same index will eventually bubble the
	// most appropriate number to the front - one that satisfies all the rules. Keep doing this for
	// each index in the update, comparing each number to all other numbers after it.
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if rules[[2]int{update[i], update[j]}] {
				update[i], update[j] = update[j], update[i]
			}
		}
	}
}

// IsUpdateValid returns true if the given update adheres to the given rules.
func IsUpdateValid(update Update, rules map[[2]int]bool) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if !rules[[2]int{update[i], update[j]}] {
				return false
			}
		}
	}
	return true
}

func ParseRules(lines []string) map[[2]int]bool {
	rules := make(map[[2]int]bool)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		before := aoc.Int(parts[0])
		after := aoc.Int(parts[1])
		rules[[2]int{before, after}] = true
	}
	return rules
}

func ParseUpdates(lines []string) []Update {
	seqs := make([]Update, 0, len(lines))
	for _, line := range lines {
		seqs = append(seqs, ParseUpdate(line))
	}
	return seqs
}

func ParseUpdate(line string) Update {
	parts := strings.Split(line, ",")
	seq := make(Update, len(parts))
	for i, part := range parts {
		seq[i] = aoc.Int(part)
	}
	return seq
}

type Update []int

func (u Update) String() string {
	s := make([]string, len(u))
	for i, val := range u {
		s[i] = fmt.Sprintf("%d", val)
	}
	return strings.Join(s, ",")
}
