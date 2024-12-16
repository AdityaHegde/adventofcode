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

func GCD(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(nums []int64) int64 {
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		ans = (nums[i] * ans) / GCD(nums[i], ans)
	}

	return ans
}

func Digits(num int) []int {
	digits := make([]int, 0)

	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}

	return digits
}
