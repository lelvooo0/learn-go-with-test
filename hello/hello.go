package main

import "fmt"

const helloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"

func main() {
	fmt.Println(Hello("world", "English"))
}

// Hello :say hello to people
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return getPrefix(language) + name
}
func getPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
