package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(reverse("Lagos Nigeria"))
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