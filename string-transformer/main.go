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
		"of":  true,
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

func snake(name string) string {
	w := strings.ToLower(name)
	return strings.ReplaceAll(w, " ", "_")
}

func main() {
	for {
		var input string
		fmt.Print("Enter a word : ")
		fmt.Scanln(&input)

		var operator string
		fmt.Print("Choose input transformations (UpperCase, LowerCase, Capitalize, Title, snake, exit): ")
		fmt.Scanln(&operator)

		if operator == "exit" {
			fmt.Println("exiting...")
			return
		}
		if operator != "UpperCase" && operator != "LowerCase" && operator != "Capitalize" && operator != "Title" && operator != "snake" && operator != "exit" {
			fmt.Println("Unrecognised")
			continue
		}

		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')

		//fmt.Print("Output 1: ")
		fmt.Println(UpperCase(name))

		//fmt.Print("Output 2: ")
		fmt.Println(LowerCase(name))

		//fmt.Print("Output 3: ")
		fmt.Println(Capitalize(name))

		//fmt.Print("Output 4: ")
		fmt.Println(Title(name))

		//fmt.Print("Output 5: ")
		fmt.Println(snake(name))
	}
}
