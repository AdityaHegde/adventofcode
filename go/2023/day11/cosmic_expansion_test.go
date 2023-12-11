package day11

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 374, partOne(utils.InputLinesFromFile("sample.txt"), 2))
	require.Equal(t, 10422930, partOne(utils.InputLinesFromFile("input.txt"), 2))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 1030, partOne(utils.InputLinesFromFile("sample.txt"), 10))
	require.Equal(t, 8410, partOne(utils.InputLinesFromFile("sample.txt"), 100))
	require.Equal(t, 699909023130, partOne(utils.InputLinesFromFile("input.txt"), 1_000_000))
}
