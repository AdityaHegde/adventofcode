package main

import (
  "fmt"
  "strings"

  utils2 "AdityaHegde/adventofcode/go/utils"
)

func main() {
  lines := utils2.InputLines()
  fmt.Println(partOne(lines))
  partTwo(lines)
}

type cpu struct {
  reg    int
  cycles int
}

func newCpu() *cpu {
  return &cpu{
    reg:    1,
    cycles: 0,
  }
}

func (c *cpu) add(amt int) {
  c.reg += amt
}

func (c *cpu) parseInstruction(instruction string, callback func(cycles int)) {
  oprn := strings.Split(instruction, " ")
  switch oprn[0] {
  case "noop":
    c.cycles++
    callback(c.cycles)
  case "addx":
    amt := int(utils2.Int64(oprn[1]))
    c.cycles++
    callback(c.cycles)
    c.cycles++
    callback(c.cycles)
    c.add(amt)
  }
}

func partOne(lines []string) int {
  res := 0
  c := newCpu()

  addToRes := func() {
    if (c.cycles+20)%40 == 0 {
      res += c.reg * c.cycles
    }
  }

  for _, line := range lines {
    c.parseInstruction(line, func(cycles int) {
      addToRes()
    })
  }

  return res
}

func partTwo(lines []string) {
  c := newCpu()
  cur := 0
  row := ""

  for _, line := range lines {
    c.parseInstruction(line, func(cycles int) {
      if (c.cycles-1)%40 == 0 {
        fmt.Print("\n")
        cur = 0
        row = ""
      }
      if cur >= c.reg-1 && cur <= c.reg+1 {
        fmt.Print("#")
        row += "#"
      } else {
        fmt.Print(".")
        row += "."
      }
      //fmt.Println(row, cur, c.cycles)
      cur++
    })
  }
}
