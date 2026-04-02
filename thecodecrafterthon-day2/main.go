package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToBinary(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)
}

func binToDecimal(bin string) (int64, error) {
	return strconv.ParseInt(bin, 2, 64)
}

func decToHex(dec int64) string {
	return strconv.FormatInt(dec, 16)
}

func decToBin(dec int64) string {
	return strconv.FormatInt(dec, 2)
}

func main() {
	for {
		var input string
		fmt.Print("Enter a word : ")
		fmt.Scanln(&input)

		var operator string
		fmt.Print("choose input bases (hex, bin, dec, exit): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("exiting..")
			return
		}
		if operator != "hex" && operator != "bin" && operator != "dec" && operator != "exit" {
			fmt.Println("invalid")
			continue
		}
		if operator == "hex" {
			i, err := hexToBinary(input)
			if err != nil {
				fmt.Println("Invalid Hex")
			}else {
			fmt.Println("Decimal: ", i)
			}
		}
		if operator == "bin" {
			i, err := binToDecimal(input)
			if err != nil {
				fmt.Println("Invalid Bin")
			}else {
				fmt.Println("Decimal: ", i)
			}
		}
		if operator == "dec" {
			i, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Invalid Decimal")
				continue
			}
			fmt.Println(strings.ToUpper(decToHex(i)))

			b, _ := strconv.ParseInt(input, 10, 64)
			fmt.Println(decToBin(b))
		}

	}

}
