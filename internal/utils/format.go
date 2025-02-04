package utils

import (
	"strings"
)

func FormatText(s string) string {
	result := formatPunctuation(s)
	result = formatArticles(result)
	result = fixSingleQuotes(result)
	result = formatPunctuation(result)
	return result
}

func formatPunctuation(s string) string {
	punctuation := []string{".", ",", "!", "?", ":", ";"}
	words := strings.Fields(s)
	result := make([]string, 0, len(words))

	for _, word := range words {
		processed := word
		
		// Handle punctuation at the end of words
		for _, p := range punctuation {
			if strings.HasSuffix(processed, p) && len(processed) > 1 {
				beforePunc := processed[:len(processed)-1]
				processed = beforePunc + p
			}
		}

		// Handle punctuation at the start of words
		for _, p := range punctuation {
			if strings.HasPrefix(processed, p) && len(processed) > 1 {
				afterPunc := processed[1:]
				processed = p + " " + afterPunc
			}
		}

		result = append(result, processed)
	}

	return strings.Join(result, " ")
}

func formatArticles(s string) string {
	silentWords := map[string]bool{
		"honest":    true,
		"heir":      true,
		"honorific": true,
		"honor":     true,
		"herb":      true,
		"hotel":     true,
		"hour":      true,
		"homage":    true,
	}

	conjunctions := map[string]bool{
		"for": true,
		"and": true,
		"nor": true,
		"but": true,
		"or":  true,
		"so":  true,
		"yet": true,
	}

	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			if conjunctions[words[i+1]] {
				continue
			}

			if silentWords[words[i+1]] || strings.ContainsRune("aeiouAEIOU", rune(words[i+1][0])) {
				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		} else if words[i] == "an" || words[i] == "An" {
			if conjunctions[words[i+1]] {
				continue
			}

			if !silentWords[words[i+1]] && !strings.ContainsRune("aeiouAEIOU", rune(words[i+1][0])) {
				if words[i] == "an" {
					words[i] = "a"
				} else {
					words[i] = "A"
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func fixSingleQuotes(s string) string {
	words := strings.Fields(s)
	result := make([]string, 0, len(words))
	inQuote := false
	quoteContent := []string{}

	for _, word := range words {
		if strings.Count(word, "'") == 1 {
			if !inQuote {
				inQuote = true
				quoteContent = []string{word}
			} else {
				inQuote = false
				quoteContent = append(quoteContent, word)
				quoted := strings.Join(quoteContent, " ")
				result = append(result, quoted)
			}
		} else if inQuote {
			quoteContent = append(quoteContent, word)
		} else {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}