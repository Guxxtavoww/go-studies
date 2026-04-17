package structs

type ExampleStruct struct {
	Field1 string
	Field2 int
}

func (es *ExampleStruct) Method1() string {
	println(es.Field2)

	return "This is Method1"
}
