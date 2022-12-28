package main

import (
	"fmt"
	"log"

	"rsc.io/quote"

	"github.com/kevinvu184/go-tutorials/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	fmt.Println(quote.Go())

	message, err := greetings.Hello("Kevin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)


    names := []string{"Gladys", "Samantha", "Darrin"}
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
