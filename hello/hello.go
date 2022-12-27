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

	message2, err2 := greetings.Hello("")
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(message2)
}
