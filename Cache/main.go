package main

import (
	"NewInMemoryCache/cache"
	"fmt"
)

func main() {

	fmt.Println("welcome to in-memory cache")

	c := cache.NewCache()
	c.Set("User1", "ajay kumar")
	c.Set("User124", "lekin kumar")
	c.Set("User56", "ragu kumar")

	fmt.Println(c.Get("User1"))

	x := (0 + 1) / 2
	fmt.Println(x)

}
