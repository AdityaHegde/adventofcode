package day9

import (
	"fmt"
	"testing"

	"AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
	fmt.Println(partOne(utils.InputLinesFromFile("sample.txt")))
	fmt.Println(partOne(utils.InputLinesFromFile("input.txt")))
}
