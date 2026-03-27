package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input1 string
		fmt.Print("Enter input1 : ")
		fmt.Scanln(&input1)
		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Enter numbers only")
			continue
		}
		var operator string
		fmt.Print("select operator (+,-,*,/, exit, help): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("exiting..")
			break
		}
		if operator == "help" {
			fmt.Println("add <a> <b> → addition")
			fmt.Println("sub <a> <b> → subtraction")
			fmt.Println("mul <a> <b> → multiplication")
			fmt.Println("sub <a> <b> → divition")
		}

		if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
			fmt.Println("Invalid operation")
			continue
		}

		var input2 string
		fmt.Print("Enter input2 : ")
		fmt.Scanln(&input2)
		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Enter numbers only")
		}
		if operator == "+" {
			fmt.Println("Total = ", num1+num2)
		} else if operator == "-" {
			fmt.Println("Total = ", num1-num2)
		} else if operator == "*" {
			fmt.Println("Total = ", num1*num2)
		} else if operator == "/" {
			if num2 == 0 {
				fmt.Println("division by zero")
			} else {
				fmt.Println("Total = ", num1/num2)
			}

		}
	}
}
