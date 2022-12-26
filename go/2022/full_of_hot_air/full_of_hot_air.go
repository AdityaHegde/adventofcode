package full_of_hot_air

var charToNum = map[uint8]int{
  '=': -2,
  '-': -1,
  '0': 0,
  '1': 1,
  '2': 2,
}
var numToChar = map[int]string{
  -2: "=",
  -1: "-",
  0:  "0",
  1:  "1",
  2:  "2",
}

func fromSNAFU(snafu string) int {
  n := 0
  for _, c := range snafu {
    n = n*5 + charToNum[uint8(c)]
  }
  return n
}

func toSNAFU(num int) string {
  snafu := ""
  for num > 0 {
    rem := num % 5
    num = num / 5
    if rem > 2 {
      num++
      rem = rem - 5
    }
    snafu = numToChar[rem] + snafu
  }
  return snafu
}

func partOne(lines []string) string {
  res := 0
  for _, line := range lines {
    res += fromSNAFU(line)
  }
  return toSNAFU(res)
}
