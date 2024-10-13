package enumgen

import "strings"

// Custom "constMaker" function transform any string into a PascalCase string constant
func constMaker(input string) string {
	// split the input string into words, anything that is not unicode letter or number, can be used as a separator
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
	})

	// convert each word to PascalCase
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	// join the words with an empty string
	return strings.Join(words, "")
}
