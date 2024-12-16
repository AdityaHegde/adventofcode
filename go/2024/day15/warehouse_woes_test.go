package day15

import (
	"fmt"
	"strings"
	"testing"

	"AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
	fmt.Println(partOne(strings.Split(`#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`, "\n")))
	fmt.Println(partOne(utils.InputLinesFromFile("sample.txt")))
	fmt.Println(partOne(utils.InputLinesFromFile("input.txt")))
}
