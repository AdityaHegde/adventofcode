package day12

import (
	"testing"

	"AdityaHegde/adventofcode/go/utils"
	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 140, partOne(utils.InputLinesFromFile("sample1.txt")))
	require.Equal(t, 772, partOne(utils.InputLinesFromFile("sample2.txt")))
	require.Equal(t, 1930, partOne(utils.InputLinesFromFile("sample3.txt")))
	require.Equal(t, 1549354, partOne(utils.InputLinesFromFile("input.txt")))
}

//func Test_partTwo(t *testing.T) {
//  require.Equal(t, 81, partTwo(utils.InputLinesFromFile("sample.txt")))
//  require.Equal(t, 1511, partTwo(utils.InputLinesFromFile("input.txt")))
//}
