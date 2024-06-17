package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return Greetings(language) + name + "!"
}

func Greetings(language string) (helloPrefix string) {
	switch language {
	case spanish:
		helloPrefix = spanishHelloPrefix
	case french:
		helloPrefix = frenchHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}

	return

}

func main() {
	name := "world"
	fmt.Println(Hello(name, "Spanish"))

	//fmt.Println(integers.Add(3, 5))
}
