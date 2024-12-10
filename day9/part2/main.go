package main

import (
	"bytes"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/deitrix/aoc/day9"
)

var input = day9.Input

func main() {
	var blocks []Block
	var id int
	for i, size := range input {
		size -= 48
		switch i % 2 {
		case 0:
			blocks = append(blocks, Block{Type: FileBlock, Size: int(size), FileID: id})
			id++
		case 1:
			blocks = append(blocks, Block{Type: FreeBlock, Size: int(size)})
		}
	}

	fileIndex := len(blocks)
	for {
		fileIndex = nextFileBlock(blocks[:fileIndex])
		if fileIndex == -1 {
			break
		}
		freeIndex := nextFreeBlock(blocks[:fileIndex], blocks[fileIndex].Size)
		if freeIndex == -1 {
			continue
		}
		// Convert the file block to a free block
		blocks[fileIndex].Type = FreeBlock

		// If the blocks are the same size, we can just change the type and copy over the file ID
		if blocks[fileIndex].Size == blocks[freeIndex].Size {
			blocks[freeIndex].Type = FileBlock
			blocks[freeIndex].FileID = blocks[fileIndex].FileID
		} else {
			// Otherwise, we need to reduce the size of the free block and insert a new file block
			// before it, making sure to update the fileIndex cursor to account for the new element.
			blocks[freeIndex].Size -= blocks[fileIndex].Size
			blocks = slices.Insert(blocks, freeIndex, Block{
				Type:   FileBlock,
				Size:   blocks[fileIndex].Size,
				FileID: blocks[fileIndex].FileID,
			})
			fileIndex++
		}
	}

	var result int
	var index int
	for _, b := range blocks {
		if b.Type == FileBlock {
			for i := range b.Size {
				result += (i + index) * b.FileID
			}
		}
		index += b.Size
	}

	fmt.Println(result)
}

func nextFileBlock(blocks []Block) int {
	for i, b := range slices.Backward(blocks) {
		if b.Type == FileBlock {
			return i
		}
	}
	return -1
}

func nextFreeBlock(blocks []Block, minSize int) int {
	for i, b := range blocks {
		if b.Type == FreeBlock && b.Size >= minSize {
			return i
		}
	}
	return -1
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
