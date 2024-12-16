package day7

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 3749, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 1708857123053, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 11387, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 189207836795655, partTwo(utils.InputLinesFromFile("input.txt")))
}
