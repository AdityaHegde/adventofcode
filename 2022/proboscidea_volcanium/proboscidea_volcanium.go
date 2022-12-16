package proboscidea_volcanium

import (
  "container/heap"
  "regexp"
  "strings"

  "AdityaHegde/adventofcode/utils"
)

type node struct {
  rate   int
  isOpen bool
  dist   map[string]int
  path   map[string]string
}

func newNode(rate int, connections []string) *node {
  n := &node{
    rate:   rate,
    isOpen: true,
    dist:   make(map[string]int),
    path:   make(map[string]string),
  }
  for _, connection := range connections {
    n.dist[connection] = 1
    n.path[connection] = connection
  }
  return n
}

type graph struct {
  nodeMap map[string]*node
}

func newGraph() *graph {
  return &graph{
    nodeMap: map[string]*node{},
  }
}

var NodeParseRegex = regexp.MustCompile(`Valve (.*) has flow rate=(\d*); tunnels? leads? to valves? (.*)$`)

func (g *graph) parseLine(line string) {
  matches := NodeParseRegex.FindStringSubmatch(line)
  node := newNode(utils.Int(matches[2]), strings.Split(matches[3], ", "))
  g.nodeMap[matches[1]] = node
}

func (g *graph) parse(lines []string) {
  for _, line := range lines {
    g.parseLine(line)
  }
  g.fillDist()
}

func (g *graph) fillDist() {
  for a, na := range g.nodeMap {
    for b, nb := range g.nodeMap {
      if b == a {
        continue
      }

      if _, ok := na.dist[b]; !ok {
        na.dist[b] = 999999
        nb.dist[a] = 999999
      }
    }
  }

  for a, na := range g.nodeMap {
    for b, nb := range g.nodeMap {
      if b == a {
        continue
      }
      for c, nc := range g.nodeMap {
        if c == a || c == b {
          continue
        }
        if na.dist[b] > na.dist[c]+nc.dist[b] {
          na.dist[b] = na.dist[c] + nc.dist[b]
          na.path[b] = c
          nb.dist[a] = na.dist[b]
          nb.path[a] = c
        }
      }
    }
  }
}

func (g *graph) fillHeap(vh *valveHeap, cur string, time int) {
  for name, n := range g.nodeMap {
    if n.isOpen && n.rate > 0 {
      weight := n.rate * (time - n.dist[cur] - 1)
      if weight < 0 {
        continue
      }
      heap.Push(vh, &valve{
        nodeName: name,
        weight:   weight,
      })
    }
  }
}

func partOne(g *graph, cur string, time int, pressure int) int {
  g.nodeMap[cur].isOpen = false

  vh := &valveHeap{}
  g.fillHeap(vh, cur, time)
  retPressure := pressure

  for time > 0 && vh.Len() > 0 {
    toV := heap.Pop(vh).(*valve)
    toN := g.nodeMap[toV.nodeName]
    newTime := time - toN.dist[cur] - 1
    //fmt.Printf("Opened %s - %d - %d\n", toV.nodeName, newTime, pressure+toN.rate*newTime)
    newPressure := partOne(
      g,
      toV.nodeName,
      newTime,
      pressure+toN.rate*newTime,
    )
    if newPressure > retPressure {
      retPressure = newPressure
    }
  }

  g.nodeMap[cur].isOpen = true
  return retPressure
}

func partTwo(g *graph, cur1, cur2 string, time1, time2 int, pressure int) int {
  g.nodeMap[cur1].isOpen = false
  if cur2 != "" {
    g.nodeMap[cur2].isOpen = false
  }

  vh1 := &valveHeap{}
  g.fillHeap(vh1, cur1, time1)
  vh2 := &valveHeap{}
  if cur2 != "" {
    g.fillHeap(vh2, cur2, time2)
  }
  retPressure := pressure

  for vh1.Len() > 0 {
    toV1 := heap.Pop(vh1).(*valve)
    v1Name := toV1.nodeName
    toN1 := g.nodeMap[v1Name]
    toN1.isOpen = false
    newTime1 := time1 - toN1.dist[cur1] - 1
    newPressure := pressure + toN1.rate*newTime1

    var toV2 *valve
    var v2Clash *valve
    v2Name := ""
    newTime2 := 0
    for vh2.Len() > 0 && toV2 == nil {
      toV2 = heap.Pop(vh2).(*valve)
      if toV2.nodeName == v1Name {
        v2Clash = toV2
        toV2 = nil
      }
    }
    if toV2 != nil {
      v2Name = toV2.nodeName
      toN2 := g.nodeMap[v2Name]
      toN2.isOpen = false
      newTime2 = time2 - toN2.dist[cur2] - 1
      newPressure += toN2.rate * newTime2
    }
    if v2Clash != nil {
      heap.Push(vh2, v2Clash)
    }

    newPressure = partTwo(
      g,
      v1Name,
      v2Name,
      newTime1,
      newTime2,
      newPressure,
    )
    if newPressure > retPressure {
      retPressure = newPressure
    }
  }

  g.nodeMap[cur1].isOpen = true
  if cur2 != "" {
    g.nodeMap[cur2].isOpen = true
  }
  return retPressure
}
