package utils

import (
  "strconv"
)

func Int(str string) int64 {
  intVal, err := strconv.ParseInt(str, 10, 64)
  if err != nil {
    panic(err)
  }
  return intVal
}
