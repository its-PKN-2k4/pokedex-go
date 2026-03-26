package main

import (
	"strings"
)

func cleanInput(text string) []string {
	sliced := strings.Split(text, " ")
	clean := []string{}
	for _, word := range sliced {
		word = strings.ToLower(strings.TrimSpace(word))
		if len(word) != 0 {
			clean = append(clean, word)
		}
	}
	return clean
}
