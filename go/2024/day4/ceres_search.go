package day4

var directions = []struct {
	dx, dy int
}{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}
var xDirections = []struct {
	x, y, dx, dy int
}{
	{1, 1, -1, -1},
	{-1, 1, 1, -1},
	{-1, -1, 1, 1},
	{1, -1, -1, 1},
}

func partOne(lines []string) int {
	count := 0

	for y, line := range lines {
		for x, c := range line {
			if c != 'X' {
				continue
			}

			for _, dir := range directions {
				if findWord(lines, "XMAS", x, y, dir.dx, dir.dy) {
					count++
				}
			}
		}
	}

	return count
}

func partTwo(lines []string) int {
	count := 0

	for y, line := range lines {
		for x, c := range line {
			if c != 'A' {
				continue
			}

			ct := 0
			for _, dir := range xDirections {
				if findWord(lines, "MAS", x+dir.x, y+dir.y, dir.dx, dir.dy) {
					ct += 1
				}
			}
			if ct == 2 {
				count += 1
			}
		}
	}

	return count
}

func findWord(lines []string, word string, x, y, dx, dy int) bool {
	c := 0
	for c < len(word) {
		if x < 0 || x >= len(lines[0]) || y < 0 || y >= len(lines) {
			return false
		}

		if lines[y][x] != word[c] {
			return false
		}

		c += 1
		x = x + dx
		y = y + dy
	}

	return true
}
