package day3

import (
	"bytes"
	"unicode"
)

func partOne(lines []string) int {
  res := 0

  for _, l := range lines {
    idx := 0
    line := []byte(l)
    for idx < len(line) {
      var mul int
      line = line[idx:]

      idx = bytes.Index(line, []byte("mul("))
      if idx == -1 {
        break
      }

      idx, mul, _ = parseMul(line, idx+4)
      res += mul
    }
  }
  return res
}

func partTwo(lines []string) int {
  res := 0
  do := true

  for _, l := range lines {
    idx := 0
    line := []byte(l)

    for idx < len(line) {
      var mul int
      line = line[idx:]

      idx = bytes.Index(line, []byte("mul("))
      if idx == -1 {
        break
      }

      doIdx := bytes.Index(line, []byte("do()"))
      dontIdx := bytes.Index(line, []byte("don't()"))
      if doIdx != -1 && (doIdx < dontIdx || dontIdx == -1) && doIdx < idx {
        idx = doIdx + 1
        do = true
        continue
      } else if dontIdx != -1 && dontIdx < idx {
        idx = dontIdx + 1
        do = false
        continue
      }

      idx, mul, _ = parseMul(line, idx+4)
      if do {
        res += mul
      }
    }
  }
  return res
}

func parseMul(line []byte, idx int) (int, int, bool) {
  idx, left, ok := parseNum(line, idx)
  if !ok {
    return idx, 0, false
  }

  if line[idx] != ',' {
    return idx, 0, false
  }
  idx += 1

  idx, right, ok := parseNum(line, idx)
  if !ok {
    return idx, 0, false
  }

  if line[idx] != ')' {
    return idx, 0, false
  }

  return idx + 1, left * right, true
}

const Num0CharCode = 48

func parseNum(line []byte, idx int) (int, int, bool) {
  if !unicode.IsNumber(rune(line[idx])) {
    return idx, 0, false
  }
  n := int(line[idx] - Num0CharCode)
  i := idx + 1
  for unicode.IsNumber(rune(line[i])) {
    n = n*10 + int(line[i]-Num0CharCode)
    i += 1
  }
  return i, n, true
}
