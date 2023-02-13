package main

import "fmt"

const italian = "Italian"
const french = "French"
const englishHelloPrefix = "Hello, "
const italianHelloPrefix = "Ciao, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case italian:
		prefix = italianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("Go!", ""))
}
