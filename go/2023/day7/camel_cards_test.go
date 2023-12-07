package day7

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 6440, partOne(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 249483956, partOne(utils.InputLinesFromFile("input.txt")))
}

func Test_partTwo(t *testing.T) {
	require.Equal(t, 5905, partTwo(utils.InputLinesFromFile("sample.txt")))
	require.Equal(t, 252137472, partTwo(utils.InputLinesFromFile("input.txt")))
}
