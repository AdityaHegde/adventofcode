package main

import (
  "AdityaHegde/adventofcode/utils"
  "fmt"
)

const ZRune = 'Z'
const ARune = 'A'
const aRune = 'a'

func charToRes(char rune) int {
  if char <= ZRune {
    return int(char - ARune + 27)
  }
  return int(char - aRune + 1)
}

func intersect(str string, chars map[rune]int, index int8, ret bool) int {
  bit := 1 << index
  full := (bit << 1) - 1

  for _, char := range str {
    _, ok := chars[char]
    if !ok {
      chars[char] = bit
    } else {
      chars[char] |= bit
      if ret && chars[char] == full {
        return charToRes(char)
      }
    }
  }

  return 0
}

func main() {
  lines := utils.InputLines()

  fmt.Println(partOne(lines))
  fmt.Println(partTwo(lines))
}

func partTwo(lines []string) int {
  res := 0

  for i := 0; i < len(lines); i += 3 {
    chars := make(map[rune]int)
    intersect(lines[i], chars, 0, false)
    intersect(lines[i+1], chars, 1, false)
    res += intersect(lines[i+2], chars, 2, true)
  }

  return res
}

func partOne(lines []string) int {
  res := 0

  for _, line := range lines {
    size := len(line) / 2

    chars := make(map[rune]int)
    intersect(line[0:size], chars, 0, false)
    res += intersect(line[size:], chars, 1, true)
  }

  return res
}
