package main

import "os"

func unwrapError[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

func main() {
	file := unwrapError(os.Create("example-file.txt"))

	size := unwrapError(file.WriteString("Olá mundo"))

	println(size)

	file.Close()
}