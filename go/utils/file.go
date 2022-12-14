package utils

import (
  "os"
  "strings"
)

func InputLines() []string {
  bytes, err := os.ReadFile(os.Args[1])
  if err != nil {
    panic(err)
  }
  input := string(bytes)
  return strings.Split(input, "\n")
}

func InputLinesFromFile(file string) []string {
  bytes, err := os.ReadFile(file)
  if err != nil {
    panic(err)
  }
  input := string(bytes)
  return strings.Split(input, "\n")
}
