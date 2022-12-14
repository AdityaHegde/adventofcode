package main

import (
  "fmt"
  "regexp"

  utils2 "AdityaHegde/adventofcode/go/utils"
)

func main() {
  lines := utils2.InputLines()
  oprn, stacks := parseStacks(lines)

  //fmt.Println(partOne(stacks, lines[oprn:]))
  fmt.Println(partTwo(stacks, lines[oprn:]))
}

func parseStacks(lines []string) (int, []*utils2.LinkedList[string]) {
  stacks := make([]*utils2.LinkedList[string], 0)

  for i, line := range lines {
    if line[1] == '1' {
      return i + 2, stacks
    }

    for j := 1; j < len(line); j += 4 {
      stackIdx := (j - 1) / 4
      if len(stacks) <= stackIdx {
        stacks = append(stacks, utils2.NewLinkedList[string]())
      }
      if line[j] == ' ' {
        continue
      }
      stacks[stackIdx].Unshift(string(line[j]))
    }
  }

  return len(lines), stacks
}

var MoveRegex = regexp.MustCompile("move (\\d*) from (\\d*) to (\\d*)")

func getMovements(line string) (int, int, int) {
  matches := MoveRegex.FindStringSubmatch(line)
  amt := int(utils2.Int64(matches[1]))
  from := int(utils2.Int64(matches[2])) - 1
  to := int(utils2.Int64(matches[3])) - 1
  return amt, from, to
}

func getOutput(stacks []*utils2.LinkedList[string]) string {
  ret := ""
  for _, stack := range stacks {
    if stack.Tail != nil {
      ret += stack.Tail.Value
    }
  }
  return ret
}

func partOne(stacks []*utils2.LinkedList[string], lines []string) string {
  for _, line := range lines {
    amt, from, to := getMovements(line)
    for i := 0; i < amt; i++ {
      stacks[to].Push(stacks[from].Pop().Value)
    }
    //fmt.Println("...")
    //for _, stack := range stacks {
    //  stack.Print()
    //}
  }

  return getOutput(stacks)
}

func partTwo(stacks []*utils2.LinkedList[string], lines []string) string {
  for _, line := range lines {
    amt, from, to := getMovements(line)
    stacks[from].ShiftTail(stacks[to], amt)
    //fmt.Println("...")
    //for _, stack := range stacks {
    //  stack.Print()
    //}
  }

  return getOutput(stacks)
}
