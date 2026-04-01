// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Ugwu Chioma Ruth]
// Squad:  [The Gophers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UpperCase(name string) string {
	return strings.ToUpper(name)
}

func LowerCase(name string) string {
	return strings.ToLower(name)
}

func Capitalize(name string) string {
	n := strings.ToLower(name)
	return strings.Title(n)
}

func Title(name string) string {
	word := map[string]bool{
		"a":   true,
		"an":  true,
		"the": true,
		"but": true,
		"or":  true,
		"for": true,
		"nor": true,
		"on":  true,
		"at":  true,
		"to":  true,
		"by":  true,
		"in":  true,
		"of":  false,
		"up":  true,
		"as":  true,
		"is":  true,
		"it":  true,
	}
	text := strings.Fields(name)
	for i, w := range text {
		if i == 0 || !word[w] {
			text[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
	}
	return strings.Join(text, " ")
}

// this handles converting words to snackcase
func snake(name string) string {
	w := strings.ToLower(name)
	return strings.ReplaceAll(w, " ", "_")
	//return strings.Join(strings.Fields(strings.ToLower(name)), "")

}

func reverse(word string) string {
	if len(word) == 0 {
		return ""
	}

	result := strings.Fields(word)
	for i := 0; i < len(result); i++ {
		result[i] = reverse(result[i][1:]) + string(result[i][0])
	}

	return strings.Join(result, " ")
}

func main() {
	for {

		fmt.Print("Enter a word : ")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		// var input string
		// fmt.Scanln(&input)

		var operator string
		fmt.Println("=== Choose input transformations ===")
		fmt.Println()
		fmt.Println("[1] UpperCase")
		fmt.Println("[2] LowerCase")
		fmt.Println("[3] Capitalize")
		fmt.Println("[4] Title")
		fmt.Println("[5] snake")
		fmt.Println("[6] reverse")
		fmt.Println("[7] exit")
		fmt.Println()
		fmt.Scanln(&operator)

		// if operator == "exit" {
		// 	fmt.Println("exiting..")
		// 	return

		// }

		switch operator {
		case "1":
			fmt.Println("Output: ", UpperCase(name))
		case "2":
			fmt.Println(LowerCase(name))
		case "3":
			fmt.Println(Capitalize(name))
		case "4":
			fmt.Println(Title(name))
		case "5":
			fmt.Println(snake(name))
		case "6":
			fmt.Println(reverse(name))
		case "7":
			fmt.Println("exiting... \n Goodbye...")
			return
		}

		if operator != "1" && operator != "2" && operator != "3" && operator != "4" && operator != "5" && operator != "6" && operator != "7" {
			fmt.Println("Unrecognised")
			// fmt.Print("Choose input transformations (UpperCase, LowerCase, Capitalize, Title, snake, reverse, exit): ")
			continue
		}

	}
}
