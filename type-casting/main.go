package main

func main() {
	var value interface{} = "Fodase"

	println(value) // Gibersh
	println(value.(string)) // Imprime o valor via casting

	result, is_valid_type := value.(int)

	if !is_valid_type {
		println("Valor do tipo não é int")
	}

	println(result)
}