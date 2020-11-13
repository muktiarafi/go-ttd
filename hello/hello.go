package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	japanHelloPrefix   = "Konnichiwa, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Japanese":
		prefix = japanHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Bambank", "Japanese"))
}
