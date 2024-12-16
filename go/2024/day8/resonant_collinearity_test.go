package day8

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 14, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 220, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 34, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 813, partTwo(utils.InputLinesFromFile("input.txt")))
}
