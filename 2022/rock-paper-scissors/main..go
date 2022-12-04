package main

import (
  "AdityaHegde/adventofcode/utils"
  "fmt"
)

const OppOffset = int('A')
const MyOffset = int('X')

func main() {
  lines := utils.InputLines()

  score := 0

  for _, line := range lines {
    opp := int(line[0]) - OppOffset
    my := int(line[2]) - MyOffset

    // 2nd part
    score += my*3 + 1
    if my == 0 {
      score += (opp + 2) % 3
    } else if my == 1 {
      score += opp
    } else {
      score += (opp + 1) % 3
    }

    // 1st part
    //score += my + 1
    //if opp == my {
    //  score += 3
    //} else if (opp+1)%3 == my {
    //  score += 6
    //}
  }

  fmt.Println(score)
}
