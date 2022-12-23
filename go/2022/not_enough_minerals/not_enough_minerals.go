package not_enough_minerals

func partOne(lines []string) int {
  res := 0
  for i, line := range lines {
    b := newBlueprint(line)
    r := b.simulate(newSimulation(24), 0) * (i + 1)
    res += r
  }
  return res
}

func partTwo(lines []string) int {
  res := 1
  for i := 0; i < len(lines) && i < 3; i++ {
    b := newBlueprint(lines[i])
    r := b.simulate(newSimulation(32), 0)
    res *= r
  }
  return res
}
