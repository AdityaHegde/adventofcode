package main

import (
  "fmt"

  "AdityaHegde/adventofcode/go/utils"
)

func main() {
  lines := utils.InputLines()
  for _, line := range lines {
    fmt.Println(line, distinctChars(line, 4), distinctChars(line, 14))
  }
}

func distinctChars(input string, distinct int) int {
  chars := make(map[uint8]int)
  count := 0

  removeFromChars := func(removed uint8) {
    chars[removed]--
    if chars[removed] == 0 {
      delete(chars, removed)
      count--
    }
  }

  addToChars := func(added uint8) {
    c, ok := chars[added]
    if !ok || c == 0 {
      chars[added] = 1
      count++
    } else {
      chars[added]++
    }
  }

  for i := 0; i < distinct; i++ {
    addToChars(input[i])
  }

  if count == distinct {
    return distinct
  }

  for i := distinct; i < len(input); i++ {
    removeFromChars(input[i-distinct])
    addToChars(input[i])
    if count == distinct {
      return i + 1
    }
  }

  return -1
}
