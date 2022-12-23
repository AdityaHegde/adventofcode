package not_enough_minerals

import (
  "math"
  "regexp"

  "AdityaHegde/adventofcode/go/utils"
)

const (
  ore      int = 0
  clay         = 1
  obsidian     = 2
  geode        = 3
)

type robot struct {
  oreReq     int
  reqType    int
  req        int
  outputType int
}

type blueprint struct {
  robots map[int]*robot
  max    []int
}

var blueprintExtract = regexp.MustCompile(`Blueprint .* (\d*) ore. .* (\d*) ore. .* (\d*) ore and (\d*) clay. .* (\d*) ore and (\d*) obsidian.`)

func newBlueprint(line string) *blueprint {
  matches := blueprintExtract.FindStringSubmatch(line)
  oreRobot := &robot{0, ore, utils.Int(matches[1]), ore}
  clayRobot := &robot{0, ore, utils.Int(matches[2]), clay}
  obsidianRobot := &robot{utils.Int(matches[3]), clay, utils.Int(matches[4]), obsidian}
  geodeRobot := &robot{utils.Int(matches[5]), obsidian, utils.Int(matches[6]), geode}
  b := &blueprint{
    map[int]*robot{
      ore:      oreRobot,
      clay:     clayRobot,
      obsidian: obsidianRobot,
      geode:    geodeRobot,
    },
    []int{0, obsidianRobot.req, geodeRobot.req, math.MaxInt},
  }
  for _, r := range b.robots {
    b.max[ore] += r.oreReq
  }
  return b
}

type simulation struct {
  min       int
  robots    []int
  resources []int
}

func newSimulation(min int) *simulation {
  return &simulation{min, []int{1, 0, 0, 0}, []int{0, 0, 0, 0}}
}

func (s *simulation) copy() *simulation {
  newRobots := make([]int, 4)
  copy(newRobots, s.robots)
  newResources := make([]int, 4)
  copy(newResources, s.resources)
  return &simulation{s.min, newRobots, newResources}
}

func (s *simulation) canAffordIn(r *robot) int {
  if s.robots[r.reqType] == 0 {
    return math.MaxInt
  }

  oreIn := float64(0)
  if s.resources[ore] < r.oreReq {
    oreIn = math.Ceil(float64(r.oreReq-s.resources[ore]) / float64(s.robots[ore]))
  }

  reqIn := float64(0)
  if s.resources[r.reqType] < r.req {
    reqIn = math.Ceil(float64(r.req-s.resources[r.reqType]) / float64(s.robots[r.reqType]))
  }
  return int(math.Max(oreIn, reqIn))
}

func (s *simulation) buildRobot(r *robot) {
  s.resources[ore] -= r.oreReq
  s.resources[r.reqType] -= r.req
  s.robots[r.outputType]++
}

func (s *simulation) progress(by int) {
  s.min -= by
  for res := range s.resources {
    s.resources[res] += s.robots[res] * by
  }
}

// A DFS where every node is when a robot is bought
func (b *blueprint) simulate(s *simulation, max int) int {
  if s.min == 0 {
    return s.resources[geode]
  }
  if s.resources[geode]+s.robots[geode]*s.min+(s.min+1)*s.min/2 <= max {
    // ignore cases where buying a geode every minute is not enough to beat max
    return max
  }

  for res := geode; res >= ore; res-- {
    if s.robots[res] == b.max[res] {
      // don't build more robots than needed
      continue
    }

    canAffordIn := s.canAffordIn(b.robots[res])
    if canAffordIn > s.min {
      // cannot afford in time
      continue
    }

    ns := s.copy()
    ns.progress(canAffordIn + 1)
    ns.buildRobot(b.robots[res])
    res := b.simulate(ns, max)
    if res > max {
      max = res
    }
  }

  return max
}
