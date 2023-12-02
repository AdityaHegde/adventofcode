package day2

import (
  "fmt"
  "regexp"
  "strings"

  "AdityaHegde/adventofcode/go/utils"
)

var (
  gameRegex  = regexp.MustCompile(`Game \d*: (.*)$`)
  roundRegex = regexp.MustCompile(`(\d*) (red|green|blue)`)
)

func parsePartOne(game string, maxCubes map[string]int) bool {
  gameMatches := gameRegex.FindStringSubmatch(game)
  if len(gameMatches) == 0 {
    panic(fmt.Sprintf("failed to parse: %s", game))
  }

  rounds := strings.Split(gameMatches[1], "; ")
  for _, round := range rounds {
    for _, cube := range roundRegex.FindAllStringSubmatch(round, -1) {
      if len(cube) == 0 {
        panic(fmt.Sprintf("failed to parse: %s", round))
      }
      ct := utils.Int(cube[1])
      if ct > maxCubes[cube[2]] {
        return false
      }
    }
  }

  return true
}

var maxCubes = map[string]int{
  "red":   12,
  "green": 13,
  "blue":  14,
}

func partOne(lines []string) int {
  res := 0
  for i, line := range lines {
    if parsePartOne(line, maxCubes) {
      res += i + 1
    }
  }
  return res
}

func parsePartTwo(game string) int {
  gameMatches := gameRegex.FindStringSubmatch(game)
  if len(gameMatches) == 0 {
    panic(fmt.Sprintf("failed to parse: %s", game))
  }

  minCubes := map[string]int{
    "red":   1,
    "green": 1,
    "blue":  1,
  }

  rounds := strings.Split(gameMatches[1], "; ")
  for _, round := range rounds {
    for _, cube := range roundRegex.FindAllStringSubmatch(round, -1) {
      if len(cube) == 0 {
        panic(fmt.Sprintf("failed to parse: %s", round))
      }
      ct := utils.Int(cube[1])
      if ct > minCubes[cube[2]] {
        minCubes[cube[2]] = ct
      }
    }
  }

  pwr := 1
  for _, ct := range minCubes {
    pwr *= ct
  }

  return pwr
}

func partTwo(lines []string) int {
  res := 0
  for _, line := range lines {
    res += parsePartTwo(line)
  }
  return res
}
