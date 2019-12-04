package statistic

import "testing"

func TestAvg(t *testing.T) {
	var numbers = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result = Avg(numbers)

	if 4.5 == result {
		t.Log("Avg PASS")
	} else {
		t.Error(result)
	}
}

func TestSTD(t *testing.T) {
	var numbers = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result = STD(numbers)

	if 3.0276503540974917 == result {
		t.Log("STD PASS")
	} else {
		t.Error(result)
	}
}
