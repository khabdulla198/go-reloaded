package core

import "strings"

func Punctuation(str []string) []string {
	punctuation := ".,!?:;"

	for i := 0; i < len(str); i++ {
		word := str[i]
		if len(word) == 0 {
			continue
		}

		// Handle groups of punctuation
		if strings.Trim(word, punctuation) == "" {
			if i > 0 {
				str[i-1] += word
				str = append(str[:i], str[i+1:]...)
				i--
			}
			continue
		}

		// Handle punctuation at the beginning of the word
		if strings.ContainsAny(string(word[0]), punctuation) && i > 0 {
			str[i-1] += string(word[0])
			str[i] = word[1:]
		}
	}

	return str
}
