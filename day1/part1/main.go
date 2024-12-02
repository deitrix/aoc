package main

import (
	"fmt"
	"slices"

	. "github.com/deitrix/aoc"
)

func main() {
	var list1, list2 []int
	for line := range Lines() {
		ids := Ints(line)
		Assert(len(ids) == 2, "Expected 2 IDs")
		list1 = append(list1, ids[0])
		list2 = append(list2, ids[1])
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var distance int
	for i := 0; i < len(list1); i++ {
		fmt.Printf("%d %d = %d\n", list1[i], list2[i], Abs(list1[i]-list2[i]))
		distance += Abs(list1[i] - list2[i])
	}

	fmt.Printf("Total distance: %d\n", distance)
}
