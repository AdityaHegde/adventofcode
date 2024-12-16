package day14

import (
	"fmt"
	"testing"

	"AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
	fmt.Println(partOne(utils.InputLinesFromFile("sample.txt"), 11, 7, 100))
	fmt.Println(partOne(utils.InputLinesFromFile("input.txt"), 101, 103, 100))
}
