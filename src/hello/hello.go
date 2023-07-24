package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	gaelic  = "Gaelic"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	gaelicPrefix  = "Dia dhuit, "
)

func getPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case gaelic:
		prefix = gaelicPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("John", ""))
}
