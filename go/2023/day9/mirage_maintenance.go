package day9

import (
  "AdityaHegde/adventofcode/go/utils"
)

func parse(line string) (int, int) {
  nums := utils.ParseInts(line)
  nl := len(nums)
  res1 := nums[nl-1]
  firsts := make([]int, 0)

  az := false
  itr := 0
  for !az && itr < nl-1 {
    nz := false

    for i := nl - 1; i > itr; i-- {
      nums[i] = nums[i] - nums[i-1]
      if nums[i] != 0 {
        nz = true
      }
    }
    firsts = append(firsts, nums[itr])
    res1 += nums[nl-1]
    itr++

    az = !nz
  }

  res2 := firsts[len(firsts)-1]
  for i := len(firsts) - 2; i >= 0; i-- {
    res2 = firsts[i] - res2
  }

  return res1, res2
}

func partOne(input []string) (int, int) {
  res1 := 0
  res2 := 0
  for _, line := range input {
    r1, r2 := parse(line)
    res1 += r1
    res2 += r2
  }
  return res1, res2
}

// 1252452206
// 1443108778
