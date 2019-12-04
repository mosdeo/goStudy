package statistic

import "math"

func Avg(numbers []float64) float64 {
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	return sum / float64(len(numbers))
}

func STD(numbers []float64) float64 {
	// 自由度
	degree_of_freedom := len(numbers) - 1

	// 離差平方和
	sum_of_deviation_square := 0.0
	for _, n := range numbers {
		sum_of_deviation_square += math.Pow(n-Avg(numbers), 2)
	}

	// 回傳相除的平方根
	return math.Sqrt(sum_of_deviation_square / float64(degree_of_freedom))
}
