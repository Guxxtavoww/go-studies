package utils

import "fmt"

func UnrwrapError[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func Must[T any](value T, err error, msgAndArgs ...any) T {
	if err != nil {
		if len(msgAndArgs) > 0 {
			panic(fmt.Sprintf("%v: %v", fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...), err))
		}

		panic(err)
	}

	return value
}
