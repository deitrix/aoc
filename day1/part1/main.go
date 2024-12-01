package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	. "github.com/deitrix/aoc"
)

func main() {
	var list1, list2 []int
	for line := range Lines() {
		fields := strings.Fields(line)
		Assert(len(fields) == 2, "Expected 2 fields")
		list1 = append(list1, Must1(strconv.Atoi(fields[0])))
		list2 = append(list2, Must1(strconv.Atoi(fields[1])))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var distance int
	for i := 0; i < len(list1); i++ {
		fmt.Printf("%d %d = %d\n", list1[i], list2[i], max(list1[i], list2[i])-min(list1[i], list2[i]))
		distance += max(list1[i], list2[i]) - min(list1[i], list2[i])
	}

	fmt.Printf("Total distance: %d\n", distance)
}
