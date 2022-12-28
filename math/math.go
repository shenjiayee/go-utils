package math

// Sum 求和
func Sum[T float64 | int64 | int](a ...T) T {
	var total T
	for _, n := range a {
		total += n
	}
	return total
}

// Avg 求平均
func Avg[T float64 | int64 | int](a ...T) T {
	return Sum(a...) / T(len(a))
}
