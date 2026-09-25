package mathutil

func Power(base float64, exp int) float64 {
	if exp == 0 {
		return 1
	}
	result := 1.0
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
