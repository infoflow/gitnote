package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Golang", "Java", "Javascript"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
