package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	greetings := []string{
		"bonjour",
		"hello",
		"namaste",
		"holla",
		"konichiwa",
	}

	rand.Seed(time.Now().UnixNano())
	l := len(greetings)
	i := rand.Intn(l - 1)
	fmt.Printf("%s\n", greetings[i])
}
