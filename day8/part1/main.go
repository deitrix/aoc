package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/deitrix/aoc"
	"github.com/deitrix/aoc/day8"
)

func main() {
	lines := slices.Collect(aoc.Lines(day8.Input))
	m := CreateMap(lines)
	PlaceAntinodes(&m)
	m.Render(os.Stdout)
	fmt.Printf("\nAntinode count: %d\n", len(m.Antinodes))
}

func PlaceAntinodes(m *Map) {
	placeAntinode := func(pos Vec2) {
		if pos.X() < 0 || pos.X() >= m.Size || pos.Y() < 0 || pos.Y() >= m.Size {
			return
		}
		m.Antinodes[pos] = struct{}{}
	}
	for _, locations := range m.Antennas {
		for i, a := range locations {
			for j, b := range locations {
				if i == j {
					continue
				}
				delta := b.Sub(a)
				placeAntinode(a.Sub(delta))
				placeAntinode(b.Add(delta))
			}
		}
	}
}

type Map struct {
	Antennas         map[byte][]Vec2
	AntennaLocations map[Vec2]byte
	Antinodes        map[Vec2]struct{}
	Size             int
}

func CreateMap(lines []string) Map {
	size := len(lines)
	antennas := make(map[byte][]Vec2)
	locations := make(map[Vec2]byte)
	for y, line := range lines {
		for x, ch := range []byte(line) {
			if ch != '.' {
				antennas[ch] = append(antennas[ch], Vec2{x, y})
				locations[Vec2{x, y}] = ch
			}
		}
	}
	return Map{
		Antennas:         antennas,
		AntennaLocations: locations,
		Size:             size,
		Antinodes:        make(map[Vec2]struct{}),
	}
}

func (m *Map) Render(w io.Writer) {
	// Map out the locations of the antennas.
	buf := new(bytes.Buffer)
	for y := range m.Size {
		for x := range m.Size {
			vec := Vec2{x, y}
			if a, ok := m.AntennaLocations[vec]; ok {
				buf.WriteByte(a)
			} else if _, ok := m.Antinodes[vec]; ok {
				buf.WriteByte('#')
			} else {
				buf.WriteByte('.')
			}
		}
		buf.WriteByte('\n')
	}
	if _, err := w.Write(buf.Bytes()); err != nil {
		panic(err)
	}
}

type Vec2 [2]int

func (v Vec2) X() int {
	return v[0]
}

func (v Vec2) Y() int {
	return v[1]
}

func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{v.X() - u.X(), v.Y() - u.Y()}
}

func (v Vec2) Add(u Vec2) Vec2 {
	return Vec2{v.X() + u.X(), v.Y() + u.Y()}
}
