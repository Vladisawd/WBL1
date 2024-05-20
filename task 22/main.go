package main

import (
	"fmt"
	"math/big"
)

// так как числа большие, сразу инициализируем их как big.int
// далее используем стандартную библиотеку "math/big"
func main() {
	var a, b, result big.Int
	var operation string
	fmt.Println("Enter first number")
	fmt.Scan(&a)
	fmt.Println("Enter second number")
	fmt.Scan(&b)
	fmt.Println("Enter operation: +, -, *, /")
	fmt.Scan(&operation)

	switch {
	case operation == "+":
		fmt.Printf("Sum result: %v\n", result.Add(&a, &b))
	case operation == "-":
		fmt.Printf("Subtraction result: %v\n", result.Sub(&a, &b))
	case operation == "*":
		fmt.Printf("Multiplicationum result: %v\n", result.Mul(&a, &b))
	case operation == "/":
		fmt.Printf("Divisionum result: %v\n", result.Div(&a, &b))
	default:
		fmt.Println("Unknown operation")
	}
}
