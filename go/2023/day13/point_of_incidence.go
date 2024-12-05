package day13

import (
	"fmt"
	"strings"
)

func transpose(input []string) []string {
	ti := make([]string, 0)
	for x := 0; x < len(input[0]); x++ {
		s := ""
		for y := 0; y < len(input); y++ {
			s += string(input[y][x])
		}
		ti = append(ti, s)
	}
	return ti
}

func checkDiff(str string, mid int) bool {
	strMid := len(str) / 2
	if mid < strMid {
		return str[:mid] == str[mid:2*mid]
	} else {
		return str[mid:] == str[:len(str)-mid]
	}
}

func checkMirror(input []string) int {
	ln := len(input[0])
	mid := ln / 2

	for x := mid; x > 0; x++ {
		//fmt.Println(x, mid, ln-x)
		for y := 0; y < len(input); y++ {
			if checkDiff(input[y], x) {
				return x
			} else if checkDiff(input[y], ln-x) {
				return ln - x
			}
		}
	}

	return 0
}

func partOne(input []string) int {
	res := 0
	for i := 0; i < len(input); i++ {
		j := i
		for ; j < len(input); j++ {
			if input[j] == "" {
				break
			}
		}

		horizontal := checkMirror(input[i:j])
		vertical := checkMirror(transpose(input[i:j]))

		fmt.Println(strings.Join(input[i:j], "\n"), "\n---", horizontal, vertical)
		if horizontal >= vertical {
			res += horizontal
		}
		if vertical >= horizontal {
			res += vertical * 100
		}
		i = j
	}
	return res
}
