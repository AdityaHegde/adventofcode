package day12

import (
	"AdityaHegde/adventofcode/go/utils"
)

func partOne(input []string) int {
	f := newField(input)

	for y, plant := range f.plants {
		for x := range plant {
			r := f.regionMap[y][x]
			r.area += 1
			r.perimeter += f.perimeter(x, y)
		}
	}

	res := 0
	for _, r := range f.regions {
		res += r.area * r.perimeter
	}
	return res
}

type field struct {
	plants    []string
	regions   []*region
	regionMap [][]*region
}

func newField(plants []string) *field {
	f := &field{
		plants:    plants,
		regions:   make([]*region, 0),
		regionMap: make([][]*region, len(plants)),
	}

	for y, plant := range plants {
		f.regionMap[y] = make([]*region, len(plant))
	}

	for y, plant := range f.plants {
		for x, p := range plant {
			r := f.regionMap[y][x]
			if r != nil {
				continue
			}
			r = f.innerPlant(x, y)
			if r != nil {
				continue
			}
			f.tracePerimeter(x-1, y, 0, uint8(p))
		}
	}

	return f
}

var dirs = []struct {
	x, y int
}{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func (f *field) perimeter(x, y int) int {
	p := f.plants[y][x]
	perimeter := 0
	for _, dir := range dirs {
		dx := x + dir.x
		dy := y + dir.y

		if !f.withinBounds(dx, dy) || f.plants[dy][dx] != p {
			perimeter += 1
		}
	}
	return perimeter
}

func (f *field) innerPlant(x, y int) *region {
	p := f.plants[y][x]
	var r *region
	for _, dir := range dirs {
		dx := x + dir.x
		dy := y + dir.y

		if !f.withinBounds(dx, dy) || f.plants[dy][dx] != p {
			continue
		}
		if f.regionMap[dy][dx] != nil {
			r = f.regionMap[dy][dx]
		}
	}
	f.regionMap[y][x] = r
	return r
}

func (f *field) tracePerimeter(x, y, dir int, plant uint8) {
	r := newRegion(plant)
	f.regions = append(f.regions, r)
	v := utils.Vector{
		X:   x,
		Y:   y,
		Dir: dir,
	}

	for {
		v.Move()
		if v.X == x && v.Y == y {
			break
		}

		rx, ry := v.RightPoint()

		if !f.withinBounds(rx, ry) || f.plants[ry][rx] != plant {
			v.RotateRight()
		} else {
			nx, ny := v.NextPoint()
			for f.withinBounds(nx, ny) && f.plants[ny][nx] == plant {
				v.RotateLeft()
				nx, ny = v.NextPoint()
			}
		}

		if f.withinBounds(rx, ry) && f.plants[ry][rx] == plant {
			f.regionMap[ry][rx] = r
		}
	}
}

func (f *field) withinBounds(x, y int) bool {
	return x >= 0 && x < len(f.plants[0]) && y >= 0 && y < len(f.plants)
}

type region struct {
	plant           string
	area, perimeter int
}

func newRegion(p uint8) *region {
	return &region{string(p), 0, 0}
}
