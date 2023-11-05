package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Shin")
	fmt.Println(message)
}
