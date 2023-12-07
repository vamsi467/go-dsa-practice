package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    fmt.Println("Enter two numbers: ")
    fmt.Scanln(&num1, &num2)

    // Perform arithmetic operations and display the results
    // Add, subtract, multiply, and divide num1 and num2
	var add = num1 + num2;
	var subtract = num1 - num2;
	var multiply = num1 * num2;
	var divide = num1 / num2;

	fmt.Println(add, subtract, multiply, divide);
}