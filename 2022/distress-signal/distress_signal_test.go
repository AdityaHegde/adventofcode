package distress_signal

import (
  "fmt"
  "testing"

  "AdityaHegde/adventofcode/utils"
)

func Test_partOne(t *testing.T) {
  fmt.Println(partOne(parse(utils.InputLinesFromFile("input_0.txt"))))
  fmt.Println(partOne(parse(utils.InputLinesFromFile("input_1.txt"))))
}

func Test_partTwo(t *testing.T) {
  fmt.Println(partTwo(parse(utils.InputLinesFromFile("input_0.txt"))))
  fmt.Println(partTwo(parse(utils.InputLinesFromFile("input_1.txt"))))
}
