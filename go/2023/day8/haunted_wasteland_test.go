package day8

import (
	"fmt"
	"testing"

	"AdityaHegde/adventofcode/go/utils"
)

func Test_partOne(t *testing.T) {
	fmt.Println(partOne(utils.InputLinesFromFile("sample1.txt")))
	fmt.Println(partOne(utils.InputLinesFromFile("sample2.txt")))
	fmt.Println(partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	fmt.Println(partTwo(utils.InputLinesFromFile("sample3.txt")))
	fmt.Println(partTwo(utils.InputLinesFromFile("input.txt")))
}
