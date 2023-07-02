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
	fmt.Printf("Addint(%d, %d) = ", a, b)
	fmt.Println(calc.Addint(a, b))
	fmt.Printf("Max(%d, %d) = ", a, b)
	fmt.Println(compare.Max(a, b))
	fmt.Println("==== main  end  ====")
}
