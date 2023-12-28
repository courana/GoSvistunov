package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func Space(bytes []byte) []byte {
	out := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}

func main() {
	input := []byte("   rfwerrf, \t   aegeergrg   .")
	fmt.Printf("Изначальная: %s\n", input)

	result := Space(input)

	fmt.Printf("Преобразованная: %s\n", result)
}
