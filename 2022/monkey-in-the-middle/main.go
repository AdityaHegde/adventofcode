package main

import (
  "fmt"
  "strings"

  "AdityaHegde/adventofcode/utils"
)

func main() {
  lines := utils.InputLines()
  ms := parseMonkeys(lines)
  //partOne(ms)
  partTwo(ms)
}

type monkey struct {
  items []int64
  oprn  uint8
  amt   int64
  div   int64
  true  int
  false int
  count int
}

func (m *monkey) parse(lines []string, start int) int {
  start++

  // parse items
  items := strings.Split(
    strings.Replace(
      strings.TrimSpace(lines[start]),
      "Starting items: ",
      "",
      1,
    ),
    ", ",
  )
  for _, item := range items {
    m.items = append(m.items, utils.Int(item))
  }

  // parse operation
  mod := strings.Split(
    strings.Replace(
      strings.TrimSpace(lines[start+1]),
      "Operation: new = old ",
      "",
      1,
    ),
    " ",
  )
  if mod[1] == "old" {
    m.oprn = '^'
  } else {
    m.oprn = mod[0][0]
    m.amt = utils.Int(mod[1])
  }

  // parse throw condition
  m.div = utils.Int(strings.Replace(
    strings.TrimSpace(lines[start+2]),
    "Test: divisible by ",
    "",
    1,
  ))
  m.true = int(utils.Int(strings.Replace(
    strings.TrimSpace(lines[start+3]),
    "If true: throw to monkey ",
    "",
    1,
  )))
  m.false = int(utils.Int(strings.Replace(
    strings.TrimSpace(lines[start+4]),
    "If false: throw to monkey ",
    "",
    1,
  )))

  return start + 6
}

func (m *monkey) modify(itemIdx int) int64 {
  m.count++
  item := m.items[itemIdx]
  switch m.oprn {
  case '+':
    item += m.amt
  case '*':
    item *= m.amt
  case '^':
    item *= item
  }
  return item
}

func (m *monkey) throw(item int64) int {
  if item%m.div == 0 {
    return m.true
  } else {
    return m.false
  }
}

func (m *monkey) string() string {
  str := ""
  for _, item := range m.items {
    str += fmt.Sprintf("%d ", item)
  }
  return str
}

func parseMonkeys(lines []string) []*monkey {
  ms := make([]*monkey, 0)
  for i := 0; i < len(lines); {
    m := &monkey{}
    i = m.parse(lines, i)
    ms = append(ms, m)
  }
  return ms
}

func partOne(ms []*monkey) {
  for round := 0; round < 20; round++ {
    for _, m := range ms {
      for i := 0; i < len(m.items); i++ {
        item := m.modify(i)
        item /= 3
        to := m.throw(item)
        ms[to].items = append(ms[to].items, item)
      }
      m.items = make([]int64, 0)
    }
    //for i, m := range ms {
    //  fmt.Printf("Monkey %d: %s\n", i, m.string())
    //}
  }

  for i, m := range ms {
    fmt.Printf("Monkey %d: %d\n", i, m.count)
  }
}

func partTwo(ms []*monkey) {
  var lcm int64 = 1
  for _, m := range ms {
    lcm *= m.div
  }

  for round := 0; round < 10000; round++ {
    for _, m := range ms {
      for i := 0; i < len(m.items); i++ {
        item := m.modify(i)
        to := m.throw(item)
        item %= lcm
        ms[to].items = append(ms[to].items, item)
      }
      m.items = make([]int64, 0)
    }
    if (round+1)%1000 == 0 {
      fmt.Println("Round", round)
      for i, m := range ms {
        fmt.Printf("Monkey %d: %d\n", i, m.count)
      }
    }
  }
}
