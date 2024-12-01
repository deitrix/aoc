package main

import (
	"fmt"
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

	var score int
	for _, v1 := range list1 {
		var count int
		for _, v2 := range list2 {
			if v1 == v2 {
				count++
			}
		}
		score += v1 * count
		fmt.Printf("score += %d * %d = %d\n", v1, count, v1*count)
	}

	fmt.Printf("Total score: %d\n", score)
}
