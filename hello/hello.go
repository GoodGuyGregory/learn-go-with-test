package main

import "fmt"

// define constants:
//  Constants should improve performance of your application
//  as it saves you creating the "Hello, " string instance every time Hello is called.
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// checks for blank string params
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
