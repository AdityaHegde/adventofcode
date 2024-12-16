package day6

import (
	"fmt"
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 41, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 5516, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	fmt.Println(partTwo(utils.InputLinesFromFile("sample.txt")))
	fmt.Println(partTwo(utils.InputLinesFromFile("input.txt")))
}
