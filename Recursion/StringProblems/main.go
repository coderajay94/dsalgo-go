package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "acbabsdazf"
	fmt.Println("string :", str)
	fmt.Println("string after ", removeA(str))
	fmt.Println("string after ", removeAWithResponse("", str))
	fmt.Println("string after using extra variable ", removeABuilder(str))
	//https://leetcode.com/problems/remove-all-occurrences-of-a-substring/description/
	fmt.Println("removed  ", removeOccurrences("daabcbaabcbc", "abc"))

}

func removeOccurrences(s string, part string) string {
	return remove(s, part, 0)
}

func remove(s string, part string, start int) string {
	//fmt.Println("string", s, "with start ", start)
	if len(s) < len(part) {
		return s
	}

	if start+len(part) <= len(s) && s[start:start+len(part)] == part {
		return remove(s[0:start]+s[start+len(part):], part, 0)
	}
	if start+len(part) > len(s) {
		return s
	}

	// Otherwise, move forward
	return remove(s, part, start+1)

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

// with a resp field
func removeAWithResponse(resp, str string) string {
	if str == "" {
		fmt.Println("this is reh resp", resp)
		return str
	}
	if str[0] == 'a' {
		return removeAWithResponse(resp, str[1:]) // skip 'a'
	}
	return string(resp+string(str[0])) + removeA(string(str[1:]))
}
