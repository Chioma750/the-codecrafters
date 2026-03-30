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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Output 1: ")
	fmt.Println(UpperCase(name))

	fmt.Print("Output 2: ")
	fmt.Println(LowerCase(name))

	fmt.Print("Output 3: ")
	fmt.Println(Capitalize(name))
}
