package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
		fmt.Print("Enter a word")
		fmt.Scanln(&input)

		var base int
		fmt.Println("Enter a number")
		fmt.Scanln(&base)

		var operator string
		fmt.Print("choose input bases (hex, bin, dec, exit): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("exiting..")
			break
		}
		if operator != "hex" && operator != "bin" && operator != "dec" && operator != "exit" {
			fmt.Println("invalid")
			continue
		}

		if operator == "hex" && operator == "bin" && operator == "dec" {
			switch operator {
			case "hex":
				a, err := strconv.ParseInt(input, base, 64)
				if err != nil {
					fmt.Println("error")
				}
				fmt.Println(strconv.FormatInt(a, 10))
			}
			continue
		}

	}
}
