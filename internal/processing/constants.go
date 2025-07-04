package processing

import (
	"strconv"
	"strings"
)

type WordFuncs func(string) string

var transforms = map[string]WordFuncs{
	"up":  strings.ToUpper,
	"low": strings.ToLower,
	"cap": capitalize,
	"hex": hex,
	"bin": bin,
}

func capitalize(text string) string {
	return strings.ToUpper(text[:1]) + strings.ToLower(text[1:])
}

func hex(text string) string {
	n, err := strconv.ParseInt(text, 16, 64)
	if err != nil {
		return text
	}
	return strconv.FormatInt(n, 10)
}

func bin(text string) string {
	n, err := strconv.ParseInt(text, 2, 64)
	if err != nil {
		return text
	}
	return strconv.FormatInt(n, 10)
}
