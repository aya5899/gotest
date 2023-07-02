package main

import (
	"fmt"

	"github.com/aya5899/gotest/pkg/calc"
	"github.com/aya5899/gotest/pkg/compare"
)

func main() {
	fmt.Println("==== main start ====")
	a := 3
	b := 5
	fmt.Printf("Addint(%d, %d) = %d\n", a, b, calc.Addint(a, b))
	fmt.Printf("Max(%d, %d) = %d\n", a, b, compare.Max(a, b))
	fmt.Println("==== main   end ====")
}
