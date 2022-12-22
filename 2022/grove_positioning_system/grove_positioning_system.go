package grove_positioning_system

import (
  "AdityaHegde/adventofcode/utils"
)

type problem struct {
  nums     []int
  idxs     []int
  orderMap map[int]int
  zeroIdx  int
  length   int
}

func newProblem(lines []string) *problem {
  p := &problem{
    nums:     make([]int, len(lines)),
    idxs:     make([]int, len(lines)),
    orderMap: map[int]int{},
    length:   len(lines),
  }
  for i, line := range lines {
    p.nums[i] = utils.Int(line)
    p.idxs[i] = i
    p.orderMap[i] = i
    if p.nums[i] == 0 {
      p.zeroIdx = i
    }
  }
  return p
}

func (p *problem) getToIndex(num, idx int) int {
  toIdx := (num + idx) % (p.length - 1)
  if toIdx < 0 {
    return p.length + toIdx - 1
  }
  return toIdx
}

func (p *problem) mix() {
  for i := 0; i < p.length; i++ {
    idx := p.orderMap[i]
    num := p.nums[idx]
    numIdx := p.idxs[idx]
    toIdx := p.getToIndex(num, idx)
    //fmt.Println(num, idx, toIdx)
    if toIdx == idx {
      continue
    } else {
      newNums := make([]int, p.length)
      newIdxs := make([]int, p.length)

      smaller := idx
      larger := toIdx
      if toIdx < idx {
        smaller = toIdx
        larger = idx
      }

      for i := 0; i < smaller; i++ {
        newNums[i] = p.nums[i]
        newIdxs[i] = p.idxs[i]
      }
      if toIdx > idx {
        for i := smaller + 1; i <= larger; i++ {
          newNums[i-1] = p.nums[i]
          newIdxs[i-1] = p.idxs[i]
          p.orderMap[p.idxs[i]]--
        }
      } else {
        for i := smaller; i < larger; i++ {
          newNums[i+1] = p.nums[i]
          newIdxs[i+1] = p.idxs[i]
          p.orderMap[p.idxs[i]]++
        }
      }
      newNums[toIdx] = num
      newIdxs[toIdx] = numIdx
      p.orderMap[p.idxs[idx]] = toIdx
      for i := larger + 1; i < p.length; i++ {
        newNums[i] = p.nums[i]
        newIdxs[i] = p.idxs[i]
      }

      p.nums = newNums
      p.idxs = newIdxs
    }
    //fmt.Println(nums, idxs)
  }
}

func (p *problem) solution() int {
  newZeroIdx := p.orderMap[p.zeroIdx]
  return p.nums[(newZeroIdx+1000)%p.length] + p.nums[(newZeroIdx+2000)%p.length] + p.nums[(newZeroIdx+3000)%p.length]
}

func partOne(lines []string) int {
  p := newProblem(lines)
  p.mix()
  return p.solution()
}

func partTwo(lines []string) int {
  p := newProblem(lines)
  for i := range p.nums {
    p.nums[i] *= 811589153
  }
  for i := 0; i < 10; i++ {
    p.mix()
  }
  return p.solution()
}
