package utils

func Sign(n int) int {
  if n == 0 {
    return 0
  } else if n < 0 {
    return -1
  } else {
    return 1
  }
}
