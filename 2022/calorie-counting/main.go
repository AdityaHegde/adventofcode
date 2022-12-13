package main

import (
  "container/heap"
  "fmt"

  "AdityaHegde/adventofcode/utils"
)

type elf struct {
  items []int64
  total int64
}

type elfHeap []*elf

func (h elfHeap) Len() int {
  return len(h)
}

func (h elfHeap) Less(i, j int) bool {
  return h[i].total > h[j].total
}

func (h elfHeap) Swap(i, j int) {
  h[i], h[j] = h[j], h[i]
}

func (h *elfHeap) Push(x interface{}) {
  *h = append(*h, x.(*elf))
}

func (h *elfHeap) Pop() interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}

func main() {
  lines := utils.InputLines()

  elves := &elfHeap{}
  e := &elf{}

  for _, line := range lines {
    if line == "" {
      if e.total < 0 {
        fmt.Println("Overflow")
      }
      heap.Push(elves, e)
      e = &elf{}
      continue
    }

    c := utils.Int(line)
    e.items = append(e.items, c)
    e.total += c
  }

  //fmt.Println(heap.Pop(elves).(*elf).total) -
  fmt.Println(heap.Pop(elves).(*elf).total + heap.Pop(elves).(*elf).total + heap.Pop(elves).(*elf).total)
}
