package main

import (
	"fmt"

	"./calc"
)

func main() {
	x := 10
	y := 4

	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Sub(x, y))
	fmt.Println(calc.Multi(x, y))
	fmt.Println(calc.Divide(x, y))

}
