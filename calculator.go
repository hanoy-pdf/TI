package main

import "fmt"

func main() {

	var Number1, Number2 float64
	var operator string

	fmt.Print("Enter the first number : ")
	fmt.Scanln(&Number1)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&Number2)

	fmt.Print("Enter the Operator (+ - * /):")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Printf("%f %s %f = %f\n", Number1, operator, Number2, Number1+Number2)

	case "-":
		fmt.Printf("%f %s %f = %f\n", Number1, operator, Number2, Number1-Number2)

	case "*":
		fmt.Printf("%f %s %f = %f\n", Number1, operator, Number2, Number1*Number2)

	case "/":
		if Number2 == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			fmt.Printf("%f %s %f = %f\n", Number1, operator, Number2, Number1/Number2)
		}
	default:
		fmt.Println("Invalid Operator")
	}

}

// Hana Dwi Lathifah
