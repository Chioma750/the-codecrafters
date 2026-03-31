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
		fmt.Print("Choose input transformations (UpperCase, LowerCase, Capitalize, Title, snake, reverse, exit): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("exiting...")
			return
		}

		switch operator {
		case "UpperCase":
			fmt.Println(UpperCase(name))
		case "LowerCase":
			fmt.Println(LowerCase(name))
		case "Capitalize":
			fmt.Println(Capitalize(name))
		case "Title":
			fmt.Println(Title(name))
		case "snake":
			fmt.Println(snake(name))
		case "reverse":
			fmt.Println(reverse(name))
		}

		if operator != "UpperCase" && operator != "LowerCase" && operator != "Capitalize" && operator != "Title" && operator != "snake" && operator != "reverse" && operator != "exit" {
			fmt.Println("Unrecognised")
			fmt.Print("Choose input transformations (UpperCase, LowerCase, Capitalize, Title, snake, reverse, exit): " )
			continue
		}

	

		//fmt.Print("Output 1: ")

		
	
		

		// fmt.Print("Output 2: ")
		// fmt.Println(LowerCase(name))

		// fmt.Print("Output 3: ")
		// fmt.Println(Capitalize(name))

		// fmt.Print("Output 4: ")
		// fmt.Println(Title(name))

		// fmt.Print("Output 5: ")
		// fmt.Println(snake(name))
	}
}
