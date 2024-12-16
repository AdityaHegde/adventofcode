package day11

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_partOne(t *testing.T) {
	require.Equal(t, 22, partOne(`125 17`, 6))
	require.Equal(t, 55312, partOne(`125 17`, 25))
	require.Equal(t, 220722, partOne(`28591 78 0 3159881 4254 524155 598 1`, 25))
	fmt.Println(partOne(`28591 78 0 3159881 4254 524155 598 1`, 75))
}

//func Test_partTwo(t *testing.T) {
//  require.Equal(t, 81, partTwo(utils.InputLinesFromFile("sample.txt")))
//  require.Equal(t, 1511, partTwo(utils.InputLinesFromFile("input.txt")))
//}
