package main

// Union type
// O ~ em Go é o operador de tipo subjacente (underlying type). Ele indica que a constraint aceita não só o tipo exato, mas qualquer tipo que tenha aquele tipo como base. Ex: se tivesse um tipo extra type MyInt int, sem o ~ o go ia gritar contigo
type ValidNumberForReduction interface {
	~int | ~float64
}

// Equivalente a <T> no typescript
func reduce_user_salaries[GenericType ValidNumberForReduction](obj map[string]GenericType) GenericType {
	var result GenericType

	for _, value := range obj {
		result += value
	}

	return result
}

// Só aceita tipos comparaveis, int com int string com string etc...
func comparable_example[GenericComparable comparable](param1 GenericComparable, param2 GenericComparable) bool {
	if param1 == param2 {
		return true
	}

	return false
}

func main() {
	salaries_map := map[string]int{
		"gugu": 10,
	}

	println(reduce_user_salaries(salaries_map))


	println(comparable_example(10, 10))
}
