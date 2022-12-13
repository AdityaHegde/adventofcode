package utils

func GetMax[Ele interface{}](array []Ele, valGetter func(e Ele) int) int {
  if len(array) == 0 {
    return -1
  }

  max := valGetter(array[0])
  maxIdx := 0

  for i := 1; i < len(array); i++ {
    val := valGetter(array[i])
    if max < val {
      max = val
      maxIdx = i
    }
  }

  return maxIdx
}
