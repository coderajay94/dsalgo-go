package main

import "fmt"

func main() {

	fmt.Println("printing sub sequenece of a string")
	sub("", "abc")
}

func sub(pr, un string) {

	if un == "" {
		fmt.Println(pr)
		return
	}

	ch := un[0]
	sub(pr+string(ch), un[1:])
	sub(pr, un[1:])
}
