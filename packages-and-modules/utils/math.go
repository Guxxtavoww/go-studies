package math

func Sum[T int | float64](a ...T) T {
	var result T

	for _, value := range a {
		result += value
	}

	return result
}
