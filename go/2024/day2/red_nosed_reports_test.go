package day2

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 2, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 483, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 5, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 528, partTwo(utils.InputLinesFromFile("input.txt"))) // 526
}
