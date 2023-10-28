package main

import "fmt"

func main() {

	var operation string
	var num1, num2 int

	fmt.Println("Calculator")
	fmt.Println("---------------------------------------")

	fmt.Println("Enter the operation: addition, subtraction, division, multiplication")
	fmt.Scanf("%s", &operation)

	fmt.Scanln()

	fmt.Println("Enter the first number: ")
	fmt.Scanf("%d", &num1)

	fmt.Scanln()

	fmt.Println("Enter the second number: ")
	fmt.Scanf("%d", &num2)

	fmt.Scanln()

	switch operation {
	case "addition":
		fmt.Println(num1 + num2)
	case "subtraction":
		fmt.Println(num1 - num2)
	case "multiplication":
		fmt.Println(num1 * num2)
	case "division":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Operation not available.")
	}

}
