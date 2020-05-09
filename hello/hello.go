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

const german = "German"
const germanHelloPrefix = "Hallo, "

func Hello(name string, language string) string {
	// checks for blank string params
	if name == "" {
		name = "World"
	}

	// refactors for language handler
	return greetingPrefix(language) + name
}

//In Go public functions start with a capital letter and private ones start with a lowercase.
func greetingPrefix(language string) (prefix string) {

	// adds switch for efficientcy
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// returns what ever prefix is set to
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
