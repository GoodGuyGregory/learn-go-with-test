package main

import "fmt"

// define constants:
//  Constants should improve performance of your application
//  as it saves you creating the "Hello, " string instance every time Hello is called.
const englishHelloPrefix = "Hello, "

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	// checks for blank string params
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
