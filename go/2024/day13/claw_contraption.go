package day13

import (
	"math"
	"regexp"

	"AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
	machines := parseInput(lines)

	res := 0

	for _, m := range machines {
		cost, ok := findDist(m.tx, m.ty, m.buttons)
		if ok {
			res += cost
		}
	}

	return res
}

var (
	buttonRegex = regexp.MustCompile(`Button [AB]: X\+(\d+), Y\+(\d+)`)
	prizeRegex  = regexp.MustCompile(`Prize: X=(\d*), Y=(\d*)`)
)

func parseInput(lines []string) []*machine {
	machines := make([]*machine, 0)

	i := 0
	for i < len(lines) {
		m := machine{
			buttons: make([]*button, 2),
		}

		buttonA := buttonRegex.FindStringSubmatch(lines[i])
		m.buttons[1] = &button{
			dx:   utils.Int(buttonA[1]),
			dy:   utils.Int(buttonA[2]),
			cost: 3,
		}

		buttonB := buttonRegex.FindStringSubmatch(lines[i+1])
		m.buttons[0] = &button{
			dx:   utils.Int(buttonB[1]),
			dy:   utils.Int(buttonB[2]),
			cost: 1,
		}

		prize := prizeRegex.FindStringSubmatch(lines[i+2])
		m.tx = utils.Int(prize[1])
		m.ty = utils.Int(prize[2])

		i += 4
		machines = append(machines, &m)
	}

	return machines
}

type machine struct {
	tx, ty  int
	buttons []*button
}

func findDist(tx, ty int, buttons []*button) (int, bool) {
	initCount := int(math.Min(float64(tx/buttons[0].dx), float64(ty/buttons[0].dy)))
	x := buttons[0].dx * initCount
	y := buttons[0].dy * initCount
	cost := initCount * buttons[0].cost

	i := initCount
	j := 0
	for i >= 0 && x >= 0 && y >= 0 && j <= 100 {
		if i <= 100 && x == tx && y == ty {
			break
		}

		cost -= buttons[0].cost
		x -= buttons[0].dx
		y -= buttons[0].dy
		i -= 1

		nx := x + buttons[1].dx
		ny := y + buttons[1].dy
		for nx <= tx && ny <= ty {
			j += 1
			cost += buttons[1].cost
			x = nx
			y = ny
			nx = x + buttons[1].dx
			ny = y + buttons[1].dy
		}
	}

	if x == tx && y == ty {
		return cost, true
	}

	return 0, false
}

type button struct {
	dx, dy, cost int
}
