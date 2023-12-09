package utils

import (
	"regexp"
	"strconv"
)

func Int64(str string) int64 {
	intVal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return intVal
}

func Int(str string) int {
	intVal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(intVal)
}

var NumbersRegex = regexp.MustCompile(`([-\d]+)`)

func ParseInts(str string) []int {
	intMatches := NumbersRegex.FindAllStringSubmatch(str, -1)
	ints := make([]int, len(intMatches))
	for i, intMatch := range intMatches {
		ints[i] = Int(intMatch[1])
	}
	return ints
}
