package beacon_exclusion_zone

import (
  "fmt"
  "math"
  "regexp"
  "strings"

  "AdityaHegde/adventofcode/utils"
)

type gridEntryType uint8

const (
  unknownEntryType gridEntryType = 0
  emptyEntryType   gridEntryType = 1
  sensorEntryType  gridEntryType = 2
  beaconEntryType  gridEntryType = 3
)

type sensor struct {
  x    int
  y    int
  dist int
}

func (s *sensor) isWithin(x, y int) bool {
  dist := utils.ManhattanDist(x, y, s.x, s.y)
  return dist <= s.dist
}

func newSensor(x, y, dist int) *sensor {
  return &sensor{
    x:    x,
    y:    y,
    dist: dist,
  }
}

type problem struct {
  grid    map[int]map[int]gridEntryType
  minX    int
  maxX    int
  minY    int
  maxY    int
  sensors []*sensor
}

func newProblem() *problem {
  return &problem{
    grid:    make(map[int]map[int]gridEntryType),
    minX:    9999999,
    maxX:    0,
    minY:    9999999,
    maxY:    0,
    sensors: make([]*sensor, 0),
  }
}

func (p *problem) updateMinMax(x, y int) {
  if x > p.maxX {
    p.maxX = x
  } else if x < p.minX {
    p.minX = x
  }
  if y > p.maxY {
    p.maxY = y
  } else if y < p.minY {
    p.minY = y
  }
}

func (p *problem) add(x, y int, entry gridEntryType) {
  if _, ok := p.grid[y]; !ok {
    p.grid[y] = make(map[int]gridEntryType)
  }
  p.grid[y][x] = entry
  p.updateMinMax(x, y)
}

func (p *problem) has(x, y int) gridEntryType {
  if _, ok := p.grid[y]; !ok {
    return unknownEntryType
  }
  if entry, ok := p.grid[y][x]; ok {
    return entry
  }
  return unknownEntryType
}

func (p *problem) withinSensor(x, y int) (bool, *sensor) {
  for _, s := range p.sensors {
    if s.isWithin(x, y) {
      return true, s
    }
  }
  return false, nil
}

var beaconExtractRegex = regexp.MustCompile(`Sensor at x=([0-9\-]*), y=([0-9\-]*): closest beacon is at x=([0-9\-]*), y=([0-9\-]*)`)

func (p *problem) parseLine(line string) {
  matches := beaconExtractRegex.FindStringSubmatch(line)
  sensorX := utils.Int(matches[1])
  sensorY := utils.Int(matches[2])
  p.add(sensorX, sensorY, sensorEntryType)
  beaconX := utils.Int(matches[3])
  beaconY := utils.Int(matches[4])
  p.add(beaconX, beaconY, beaconEntryType)

  dist := utils.ManhattanDist(sensorX, sensorY, beaconX, beaconY)
  p.updateMinMax(sensorX+dist, sensorY)
  p.updateMinMax(sensorX-dist, sensorY)
  p.updateMinMax(sensorX, sensorY+dist)
  p.updateMinMax(sensorX, sensorY-dist)

  p.sensors = append(p.sensors, newSensor(sensorX, sensorY, dist))
}

func (p *problem) parse(lines []string) {
  for _, line := range lines {
    p.parseLine(line)
  }
}

func partOne(p *problem, col int) int {
  res := 0

  for x := p.minX; x <= p.maxX; x++ {
    if ok, _ := p.withinSensor(x, col); ok && p.has(x, col) == unknownEntryType {
      res++
    }
  }

  return res
}

func partTwo(p *problem) int {
  points := make(map[string]bool)
  add := func(x, y int) {
    if ok, _ := p.withinSensor(x, y); ok {
      return
    }
    key := fmt.Sprintf("%d_%d", x, y)
    points[key] = true
  }

  for i := 0; i < len(p.sensors); i++ {
    for j := i + 1; j < len(p.sensors); j++ {
      s1 := p.sensors[i]
      s2 := p.sensors[j]
      if s1.isWithin(s2.x, s2.y) {
        continue
      }
      if utils.ManhattanDist(s1.x, s1.y, s2.x, s2.y) != s1.dist+s2.dist+2 {
        continue
      }
      dx := int(math.Abs(float64(s2.x - s1.x)))
      xSign := utils.Sign(s2.x - s1.x)
      dy := int(math.Abs(float64(s2.y - s1.y)))
      ySign := utils.Sign(s2.y - s1.y)
      if dx < dy {
        for x := 0; x <= dx; x++ {
          add(s1.x+x*xSign, s1.y+(s1.dist-x+1)*ySign)
        }
      } else {
        for y := 0; y <= dx; y++ {
          add(s1.x+(s1.dist-y+1)*xSign, s1.y+y*ySign)
        }
      }
    }
  }

  for key, _ := range points {
    coords := strings.Split(key, "_")
    x := utils.Int(coords[0])
    y := utils.Int(coords[1])
    return x*4000000 + y
  }
  return 0
}
