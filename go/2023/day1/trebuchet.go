package day1

import (
  "slices"
  "strings"

  "AdityaHegde/adventofcode/go/utils"
)

func partOne(lines []string) int {
  res := 0
  for _, line := range lines {
    val := ""
    for i := 0; i < len(line); i++ {
      if line[i] >= '0' && line[i] <= '9' {
        val += string(line[i])
        break
      }
    }

    for i := len(line) - 1; i >= 0; i-- {
      if line[i] >= '0' && line[i] <= '9' {
        val += string(line[i])
        break
      }
    }

    res += utils.Int(val)
  }

  return res
}

var nums = []struct {
  word, num string
}{
  {"one", "1"},
  {"two", "2"},
  {"three", "3"},
  {"four", "4"},
  {"five", "5"},
  {"six", "6"},
  {"seven", "7"},
  {"eight", "8"},
  {"nine", "9"},
}

func partTwo(lines []string) int {
  t := utils.NewTrie[string]()
  for _, n := range nums {
    t.Add(n.word, 0, n.num)
    rw := strings.Split(n.word, "")
    slices.Reverse(rw)
    t.Add(strings.Join(rw, ""), 0, n.num)
  }

  res := 0
  for _, line := range lines {
    val := ""
    for i := 0; i < len(line); i++ {
      if line[i] >= '0' && line[i] <= '9' {
        val += string(line[i])
        break
      } else if n, ok := t.Get(line, i); ok {
        val += n.Val
        break
      }
    }

    for i := len(line) - 1; i >= 0; i-- {
      if line[i] >= '0' && line[i] <= '9' {
        val += string(line[i])
        break
      } else if n, ok := t.GetReverse(line, i); ok {
        val += n.Val
        break
      }
    }

    res += utils.Int(val)
  }

  return res
}
