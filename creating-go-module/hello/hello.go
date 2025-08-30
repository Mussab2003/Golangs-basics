package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	name := [] string{
		"Mussab",
		"Haider",
		"Arqam",
	}
	message, err := greetings.Hellos(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

}