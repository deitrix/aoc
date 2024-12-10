package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/deitrix/aoc/day9"
)

var input = day9.Example

func main() {
	var blocks []Block
	var id int
	for i, size := range input {
		size -= 48
		switch i % 2 {
		case 0:
			// file block
			blocks = append(blocks, Block{Type: FileBlock, Size: int(size), FileID: id})
			id++
		case 1:
			// free block
			blocks = append(blocks, Block{Type: FreeBlock, Size: int(size)})
		}
	}
	Render(os.Stdout, blocks)
	fmt.Println()
	fileIndex := len(blocks)
	freeIndex := 0
	for freeIndex < fileIndex {
		if fileIndex == 0 {
			break
		}
		for i, b := range slices.Backward(blocks[:fileIndex]) {
			if b.Type == FileBlock {
				fileIndex = i
				break
			}
		}
		for i, b := range blocks[freeIndex:] {
			if b.Type == FreeBlock {

			}
		}
	}
}

func Render(w io.Writer, blocks []Block) {
	buf := new(bytes.Buffer)
	for _, block := range blocks {
		switch block.Type {
		case FileBlock:
			buf.WriteString(strings.Repeat(string(byte(block.FileID)+48), block.Size))
		case FreeBlock:
			buf.WriteString(strings.Repeat(".", block.Size))
		}
	}
	if _, err := buf.WriteTo(w); err != nil {
		panic(err)
	}
}

type Block struct {
	Type   BlockType
	Size   int
	FileID int // only used if Type == FileBlock
}

type BlockType uint

const (
	FreeBlock BlockType = iota
	FileBlock
)
