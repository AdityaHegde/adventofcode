package utils

import (
  "strconv"
)

func Int64(str string) int64 {
  intVal, err := strconv.ParseInt(str, 10, 64)
  if err != nil {
    panic(err)
  }
  return intVal
}

func Int(str string) int {
  intVal, err := strconv.ParseInt(str, 10, 64)
  if err != nil {
    panic(err)
  }
  return int(intVal)
}
