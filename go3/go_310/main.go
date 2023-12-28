package ch3

import (
	"bytes"
	"fmt"
	"strings"
)

func Ex10() {
	println("Enter the string: ")
	var str string
	fmt.Scan(&str)
	println("Done:\n", comma_byte(str))
}

func comma_byte(s string) string {
	b := &bytes.Buffer{}
	n := len(s) % 3
	if n == 0 {
		n = 3
	}
	b.WriteString(s[:n])
	for i := n; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}

	return b.String()
}

func Ex11() {
	println("Enter the float number: ")
	var str string
	fmt.Scan(&str)
	println("Done:\n", comma_float(str))
}

func comma_float(s string) string {
	b := &bytes.Buffer{}

	signIndex := 0
	if s[0] == '-' || s[0] == '+' {
		b.WriteByte(s[0])
		signIndex = 1
	} // есть ли знак

	commaIndex := strings.Index(s, ".")
	if commaIndex == -1 {
		commaIndex = len(s)
	} // индекс точки

	text := s[signIndex:commaIndex]

	n := len(text) % 3
	if n > 0 {

		b.Write([]byte(text[:n]))
		if len(text) > n {
			b.WriteString(",")
		}
	} // до точки

	for i, c := range text[n:] {
		if i%3 == 0 && i != 0 {
			b.WriteRune(',')
		}
		b.WriteRune(c)
	} // после точки

	b.WriteString(s[commaIndex:])

	return b.String()
}
