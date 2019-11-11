package myLeetCode

import "math"

func reverse(x int) int {

	var abs_x int = int(math.Abs(float64(x)))
	var digis []int

	//分解十進位位數
	for i := 0; int(math.Pow10(i)) <= abs_x; i++ {
		digis = append(digis, (abs_x%int(math.Pow10(i+1)))/int(math.Pow10(i)))
	}

	//計算結果
	var reversed_x = 0
	for i := 1; i <= len(digis); i++ {
		reversed_x += digis[len(digis)-i] * int(math.Pow10(i-1))
	}

	if math.MinInt32 <= reversed_x && reversed_x <= math.MaxInt32 {
		if x >= 0 {
			return reversed_x
		} else {
			return -reversed_x
		}
	} else {
		return 0
	}
}
