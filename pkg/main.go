package main

import (
	"fmt"

	"github.com/aya5899/gotest/pkg/calc"
	"github.com/aya5899/gotest/pkg/compare"
)

func main() {
	fmt.Println(calc.Addint(3, 5))
	fmt.Println(compare.Max(3, 5))
}
