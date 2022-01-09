package main

import "fmt"

const (
	EnglishGreeting    = "Hello, "
	SpanishGreeting    = "Hola, "
	PortugueseGreeting = "Olá, "
	FrenchGreeting     = "Bonjour, "
)

func Hello(word string, language string) string {

	var greeting string

	switch language {
	case "English":
		greeting = EnglishGreeting
	case "Spanish":
		greeting = SpanishGreeting
	case "Portuguese":
		greeting = PortugueseGreeting
	case "French":
		greeting = FrenchGreeting
	}

	if word == "" {
		return greeting + "World"
	}
	return greeting + word
}

func main() {
	fmt.Println(Hello("world", "English"))
	fmt.Println(Hello("mundo", "Spanish"))
	fmt.Println(Hello("mundo", "Portuguese"))
	fmt.Println(Hello("monde", "French"))
}
