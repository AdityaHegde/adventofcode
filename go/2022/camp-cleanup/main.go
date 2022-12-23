package main

import (
  "fmt"
  "strings"

  utils2 "AdityaHegde/adventofcode/go/utils"
)

func main() {
  lines := utils2.InputLines()
  fmt.Println(getOverlaps(lines, completeOverlapOne))
  fmt.Println(getOverlaps(lines, completeOverlapTwo))
}

func completeOverlapOne(l1 int64, r1 int64, l2 int64, r2 int64) bool {
  if l1 > l2 {
    return r1 <= r2
  } else if l1 == l2 {
    return true
  } else {
    return r2 <= r1
  }
}

func completeOverlapTwo(l1 int64, r1 int64, l2 int64, r2 int64) bool {
  return l1 <= r2 && l2 <= r1
}

func getOverlaps(lines []string, overlapCheck func(int64, int64, int64, int64) bool) int {
  res := 0

  for _, line := range lines {
    elves := strings.Split(line, ",")
    elfRanges := make([][]int64, len(elves))
    for i, elf := range elves {
      ranges := strings.Split(elf, "-")
      elfRanges[i] = []int64{utils2.Int64(ranges[0]), utils2.Int64(ranges[1])}
    }

    if overlapCheck(elfRanges[0][0], elfRanges[0][1], elfRanges[1][0], elfRanges[1][1]) {
      res++
    }
  }

  return res
}
