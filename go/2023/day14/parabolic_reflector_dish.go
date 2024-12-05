package day14

import "AdityaHegde/adventofcode/go/utils"

func partOne(input []string) int {
  res := 0
  for x := 0; x < len(input[0]); x++ {
    e := 0
    for y := 0; y < len(input); y++ {
      if input[y][x] == '.' {
        continue
      }
      if input[y][x] == '#' {
        e = y + 1
        continue
      }

      res += len(input) - e
      e++
    }
  }
  return res
}

type rock struct {
  x, y int
}

type platform struct {
  roundRocks *utils.TwoDGrid[*rock]
  cubeRocks  *utils.TwoDGrid[*rock]
}

func parse(input []string) *platform {
  p := &platform{
    roundRocks: utils.NewTwoDGrid[*rock](),
    cubeRocks:  utils.NewTwoDGrid[*rock](),
  }

  for y, line := range input {
    for x, c := range line {
      if c == '.' {
        continue
      }
      if c == '#' {
        p.cubeRocks.Set(x, y, &rock{x, y})
      } else {
        p.roundRocks.Set(x, y, &rock{x, y})
      }
    }
  }
  return p
}

func (p *platform) tiltHorizontal() {

}
