package stringsUtil

import (
	"unicode"
	"strings"
	"regexp"
)


// ToSnake convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

// 文章をスネークケースに
func ToSentenceSnake(in string) string {
	chunk := SplitToWordsWithSep(in)
	for key, value := range chunk {
		chunk[key] = ToSnake(value)
	}
	return strings.Join(chunk, "")
}

// Split to Words With Separator
func SplitToWordsWithSep(in string) []string {
	return regexp.MustCompile("[a-zA-Z]+|[^a-zA-Z]+").FindAllString(in, -1)
}


