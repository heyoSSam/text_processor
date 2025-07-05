package processing

import (
	"regexp"
	"strings"
)

func FixPunctuation(text string) string {
	// clear spaces from quotes, then add from both sides
	re := regexp.MustCompile(`'\s+`)
	text = re.ReplaceAllString(text, "'")

	re = regexp.MustCompile(`\s+'`)
	text = re.ReplaceAllString(text, "'")

	re = regexp.MustCompile(`([^ ])'(\w+)'([^ ])`)
	text = re.ReplaceAllString(text, `$1 '$2' $3`)

	re = regexp.MustCompile(`([^ ])'(\w+)'`)
	text = re.ReplaceAllString(text, `$1 '$2'`)

	// work with punctuation marks
	re = regexp.MustCompile(`\s+([.,!?;:])`)
	text = re.ReplaceAllString(text, "$1")

	re = regexp.MustCompile(`([.,!?;:])([^\s'⟦])`)
	text = re.ReplaceAllString(text, "$1 $2")

	text = fixArticles(text)

	return text
}

func fixArticles(text string) string {
	re := regexp.MustCompile(`\b(a|an)\s+([a-zA-Z]+)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Split(match, " ")
		if len(parts) != 2 {
			return match
		}
		word := parts[1]
		first := strings.ToLower(string(word[0]))

		if strings.Contains("aeiouh", first) {
			return "an " + word
		}
		return "a " + word
	})
}
