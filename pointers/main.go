package main

import "fmt"

// Se não usar o ponteiro o Go vai criar uma copia do valor da variavel passada e passar para a função
func copied_params_ex(a, b int) int {
	return a + b
}

// Agora dá para manipular os valores originais dentro da função
func refed_params_ex(a, b *int) int {
	// *a = 50 -> Vai atualizar no escopo de cima

	return *a + *b
}

// refed_params_ex(&var1, &var2)

// stack retorna um int por valor — `x` é alocado na stack e copiado no retorno.
// O compilador não precisa mover `x` para a heap porque nenhuma referência a ele escapa.
func stack() int {
	x := 10
	return x
}

// heap retorna um ponteiro para `x` — como a referência escapa da função,
// o compilador (escape analysis) move `x` para a heap automaticamente.
func heap() *int {
	x := 10
	return &x
}

// alter_value_with_pointer demonstra leitura e escrita via ponteiro.
func alter_value_with_pointer() {
	a := 10

	// `pointer` armazena o endereço de `a`; `*pointer` desreferencia e acessa o valor
	var pointer *int = &a
	*pointer = 20 // equivale a: a = 20

	// `b` é outro ponteiro para o mesmo endereço de `a`
	b := &a
	*b = 300 // equivale a: a = 300

	// imprime 300 — o valor atual no endereço apontado por `b`
	fmt.Println(*b)
}

func main() {
	// `:=` infere o tipo (int), aloca `value` na stack e associa o nome ao endereço onde 10 está armazenado
	value := 10

	// `&value` captura o endereço; isso pode fazer o compilador promover `value` para a heap (escape analysis)
	var pointer *int = &value

	// imprime o endereço de memória de `value`, não o valor em si
	fmt.Println(pointer)

	alter_value_with_pointer()
}
