package main

import (
	"errors"
	"fmt"
)

const MAX_SUM int = 50

func sum(first int, secound int) (int, bool) {
	result := first + secound

	if result >= MAX_SUM {
		return result, true
	}

	return result, false
}

func assert_with_error(number int) (int, error) {
	if number == MAX_SUM {
		return 0, errors.New("Número e mt grande")
	}

	return number, nil
}

func spread_sum(numbers ...int) int {
	var sum int

	for index, number := range numbers {
		fmt.Print(index)

		sum += number
	}

	return sum
}

func closure_func() {
	total := func() int {
		return spread_sum(1, 2, 3, 4, 5) * 2
	}()

	fmt.Print(total)
}

func main() {
	res, is_max := sum(10, 40)

	if is_max {
		fmt.Println("IS MAX")
	}

	fmt.Printf("%d\n", res)

	_, error := assert_with_error(res)

	if error != nil {
		fmt.Printf("%s\n", error.Error())
	}
}
