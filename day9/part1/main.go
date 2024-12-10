package main

import (
	"fmt"
	"slices"

	"github.com/deitrix/aoc/day9"
)

var input = day9.Input

func main() {
	var diskMap []int
	var id int
	for i := 0; i < len(input); i += 2 {
		diskMap = append(diskMap, slices.Repeat([]int{id}, int(input[i]-48))...)
		if len(input) > i+1 {
			diskMap = append(diskMap, slices.Repeat([]int{-1}, int(input[i+1]-48))...)
		}
		id++
	}

	freeIndex := 0
	dataIndex := len(diskMap) - 1
	shift := func() bool {
		for diskMap[freeIndex] != -1 {
			freeIndex++
		}
		for diskMap[dataIndex] == -1 {
			dataIndex--
		}
		return freeIndex < dataIndex
	}

	for shift() {
		diskMap[freeIndex], diskMap[dataIndex] = diskMap[dataIndex], diskMap[freeIndex]
	}

	var result int
	for i, id := range diskMap {
		if id == -1 {
			continue
		}
		result += id * i
	}
	fmt.Println(result)
}
