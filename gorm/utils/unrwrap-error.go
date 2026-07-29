package utils

func UnrwrapError[T any](value T, err error) T {
	if err != nil {
		panic("Error")
	}

	return value
}