package distress_signal

import (
  "fmt"
  "sort"
)

type packet struct {
  isNum  bool
  num    int
  arr    []*packet
  parent *packet
}

func newNumPacket(num int, parent *packet) *packet {
  return &packet{
    isNum:  true,
    num:    num,
    parent: parent,
  }
}

func newPacketGroup(parent *packet) *packet {
  return &packet{
    isNum:  false,
    arr:    make([]*packet, 0),
    parent: parent,
  }
}

func (p *packet) wrap() *packet {
  newP := newPacketGroup(p.parent)
  newP.arr = append(newP.arr, p)
  return newP
}

func (p *packet) parse(line string) {
  // strip begining and last brackets
  line = line[1 : len(line)-1]

  cur := p
  var n *packet
  for _, ch := range line {
    switch ch {
    case '[':
      next := newPacketGroup(cur)
      cur.arr = append(cur.arr, next)
      cur = next
    case ']':
      if n != nil {
        cur.arr = append(cur.arr, n)
        n = nil
      }
      cur = cur.parent
    case ',':
      if n != nil {
        cur.arr = append(cur.arr, n)
        n = nil
      }
    default:
      if n != nil {
        n.num = n.num*10 + int(ch-'0')
      } else {
        n = newNumPacket(int(ch-'0'), cur)
      }
    }
  }
  if n != nil {
    cur.arr = append(cur.arr, n)
    n = nil
  }
}

func (p *packet) string() string {
  if p.isNum {
    return fmt.Sprintf("%d", p.num)
  } else {
    str := "["
    for i, cp := range p.arr {
      str += cp.string()
      if i < len(p.arr)-1 {
        str += ","
      }
    }
    return str + "]"
  }
}

func (p *packet) inOrder(right *packet) int {
  if p.isNum && right.isNum {
    if p.num > right.num {
      return -1
    } else if p.num < right.num {
      return 1
    }
    return 0
  }

  i := 0
  j := 0

  for i < len(p.arr) && j < len(right.arr) {
    lchild := p.arr[i]
    rchild := right.arr[i]
    i++
    j++

    if lchild.isNum != rchild.isNum {
      if lchild.isNum {
        lchild = lchild.wrap()
      }
      if rchild.isNum {
        rchild = rchild.wrap()
      }
    }

    cmp := lchild.inOrder(rchild)
    if cmp != 0 {
      return cmp
    }
  }

  if i < len(p.arr) {
    return -1
  }
  if j < len(right.arr) {
    return 1
  }
  return 0
}

func parse(lines []string) []*packet {
  ps := make([]*packet, 0)
  for i := 0; i < len(lines); i += 3 {
    left := newPacketGroup(nil)
    left.parse(lines[i])
    right := newPacketGroup(nil)
    right.parse(lines[i+1])
    ps = append(ps, left, right)
  }
  return ps
}

func partOne(ps []*packet) int {
  res := 0
  for i := 0; i < len(ps); i += 2 {
    inOrder := ps[i].inOrder(ps[i+1])
    if inOrder >= 0 {
      res += i/2 + 1
    }
  }
  return res
}

func partTwo(ps []*packet) int {
  two := newPacketGroup(nil)
  two.parse("[[2]]")
  six := newPacketGroup(nil)
  six.parse("[[6]]")
  ps = append(ps, two, six)

  sort.Slice(ps, func(i, j int) bool {
    return ps[i].inOrder(ps[j]) == 1
  })

  res := 1
  for i, p := range ps {
    if len(p.arr) == 1 && !p.arr[0].isNum && len(p.arr[0].arr) == 1 && (p.arr[0].arr[0].num == 2 || p.arr[0].arr[0].num == 6) {
      res *= i + 1
    }
  }
  return res
}
