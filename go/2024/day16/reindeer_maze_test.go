package day16

import (
	"fmt"
	"testing"

	"AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
	fmt.Println(partOne(utils.InputLinesFromFile("sample1.txt")))
	fmt.Println(partOne(utils.InputLinesFromFile("sample2.txt")))
	fmt.Println(partOne(utils.InputLinesFromFile("input.txt"))) // 134588
}
