package processing

import "regexp"

var tokenRegex = regexp.MustCompile(`\(\w+(?:,\s*\d+)?\)|[.,!?;:]|'|\w+`)

func Tokenize(text string) []string {
	return tokenRegex.FindAllString(text, -1)
}
