package main

const (
	spanish            = "spanish"
	french             = "french"
	helloPrefixEnglish = "Hello, "
	helloPrefixSpanish = "Hola, "
	helloPrefixFrench  = "Bonjour, "
)

func HelloWorld() string {
	return "Hello World!"
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		return helloPrefixSpanish + name
	case french:
		return helloPrefixFrench + name
	default:
		return helloPrefixEnglish + name
	}
}
