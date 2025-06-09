package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "acbabsdazf"
	fmt.Println("string :", str)
	fmt.Println("string after ", removeA(str))
	fmt.Println("string after using builder ", removeABuilder(str))

}

// efficient solution with log(n)
func removeABuilder(str string) string {
	var builder strings.Builder
	for _, value := range str {
		if value != 'a' {
			builder.WriteRune(value)
		}
	}
	return builder.String()
}

// inefficient solution with log(n2)
func removeA(str string) string {
	fmt.Println("printing the string", str)
	if str == "" {
		return str
	}
	if str[0] == 'a' {
		return removeA(str[1:]) // skip 'a'
	}
	return string(str[0]) + removeA(string(str[1:]))
}
