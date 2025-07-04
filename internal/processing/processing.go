package processing

import (
	"regexp"
	"strconv"
)

func Process(tokens []string) []string {
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		cmd, n, ok := ParseCommand(token)
		if !ok {
			result = append(result, token)
			continue
		}

		transform, exists := transforms[cmd]
		if !exists {
			continue
		}

		count := 0
		for j := len(result) - 1; j >= 0 && count < n; j-- {
			if isWord(result[j]) {
				result[j] = transform(result[j])
				count++
			}
		}
	}

	return result
}

var commandRegex = regexp.MustCompile(`^\((\w+)(?:,\s*(\d+))?\)$`)

func ParseCommand(token string) (cmd string, n int, ok bool) {
	match := commandRegex.FindStringSubmatch(token)
	if len(match) == 0 {
		return "", 0, false
	}
	cmd = match[1]
	n = 1 // по умолчанию
	if match[2] != "" {
		var err error
		n, err = strconv.Atoi(match[2])
		if err != nil {
			return "", 0, false
		}
	}
	return cmd, n, true
}

var wordRegex = regexp.MustCompile(`^\w+$`)

func isWord(s string) bool {
	return wordRegex.MatchString(s)
}
