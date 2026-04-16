package main

// loads all the packages in the current module, including the ones in the vendor directory
import "packages-and-modules/utils"

func main() {
	println(math.Sum(1, 2, 3))
}
