package day3

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 161, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 179571322, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 48, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 103811193, partTwo(utils.InputLinesFromFile("input.txt")))
}
