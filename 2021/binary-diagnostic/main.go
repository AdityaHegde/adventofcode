package main

import (
  "AdityaHegde/adventofcode/utils"
  "fmt"
)

func main() {
  lines := utils.InputLines()

  zeros := make([]int64, len(lines[0]))
  ones := make([]int64, len(lines[0]))

  for _, line := range lines {
    for i, bit := range line {
      if bit == '0' {
        zeros[i]++
      } else {
        ones[i]++
      }
    }
  }

  gr := 0
  er := 0
  mul := 1
  for i := len(zeros) - 1; i >= 0; i-- {
    if zeros[i] > ones[i] {
      er += mul
    } else {
      gr += mul
    }
    mul = mul << 1
  }

  fmt.Println(gr * er)
}
