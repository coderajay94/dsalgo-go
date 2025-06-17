package main

import (
	"fmt"
)

func main() {
	permutations("", "abc")
}

func permutations(prefix, remaining string) {
	if len(remaining) == 0 {
		fmt.Println(prefix)
		return
	}

	for i := 0; i < len(remaining); i++ {
		ch := string(remaining[i])
		rest := remaining[:i] + remaining[i+1:]
		fmt.Println("printing:", prefix, remaining, ch, rest)
		permutations(prefix+ch, rest)
	}
}
