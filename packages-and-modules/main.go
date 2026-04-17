package main

// loads all the packages in the current module, including the ones in the vendor directory
// also aliass the packages to avoid name conflicts
import (
	structs "packages-and-modules/structs"
	math "packages-and-modules/utils"
)

func main() {
	println(math.Sum(1, 2, 3))

	ex := &structs.ExampleStruct{
		Field1: "fodase",
		Field2: 10101,
	}

	ex.Method1()
}
