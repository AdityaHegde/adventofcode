package main

import (
  "fmt"
  "strings"

  "AdityaHegde/adventofcode/utils"
)

type file struct {
  name string
  size int64
}

func (f *file) print(level int) {
  indent := strings.Repeat(" ", level*2)
  fmt.Printf("%s- %s (flie, size=%d)\n", indent, f.name, f.size)
}

type directory struct {
  name   string
  files  map[string]*file
  dirs   map[string]*directory
  size   int64
  parent *directory
}

func newDirectory(name string, parent *directory) *directory {
  return &directory{
    name:   name,
    files:  make(map[string]*file, 0),
    dirs:   make(map[string]*directory, 0),
    size:   0,
    parent: parent,
  }
}

func (d *directory) print(level int) {
  indent := strings.Repeat(" ", level*2)
  fmt.Printf("%s- %s (dir, size=%d)\n", indent, d.name, d.size)
  for _, dir := range d.dirs {
    dir.print(level + 1)
  }
  for _, file := range d.files {
    file.print(level + 1)
  }
}

func main() {
  lines := utils.InputLines()
  root := parseDirectories(lines)
  fmt.Println("RootSize", root.size)
  fmt.Println(partOne(root))
  fmt.Println(partTwo(root, root.size, root.size))
}

func parseDirectories(lines []string) *directory {
  root := newDirectory("/", nil)
  cur := root

  for i := 1; i < len(lines); i++ {
    if strings.HasPrefix(lines[i], "$") {
      cmd := strings.Split(lines[i][2:], " ")
      switch cmd[0] {
      case "ls":
        continue
      case "cd":
        if cmd[1] == ".." {
          // update parent after finishing parsing child
          cur.parent.size += cur.size
          cur = cur.parent
        } else {
          cur = cur.dirs[cmd[1]]
        }
      }
    } else {
      entry := strings.Split(lines[i], " ")
      if entry[0] == "dir" {
        cur.dirs[entry[1]] = newDirectory(entry[1], cur)
      } else {
        size := utils.Int64(entry[0])
        cur.files[entry[1]] = &file{
          name: entry[1],
          size: size,
        }
        cur.size += size
      }
    }
  }

  for cur != root {
    // update parent after finishing parsing everything
    cur.parent.size += cur.size
    cur = cur.parent
  }

  return root
}

const PartOneLimit = 100000

func partOne(root *directory) int64 {
  var res int64 = 0
  if root.size <= PartOneLimit {
    res += root.size
  }

  for _, dir := range root.dirs {
    res += partOne(dir)
  }

  return res
}

const TotalDiskSpace = 70000000
const DiskSpaceNeeded = 30000000

func partTwo(root *directory, rootSize int64, curSize int64) int64 {
  needed := DiskSpaceNeeded - (TotalDiskSpace - rootSize)
  if root.size <= curSize && root.size >= needed {
    curSize = root.size
  }

  for _, dir := range root.dirs {
    curSize = partTwo(dir, rootSize, curSize)
  }

  return curSize
}
