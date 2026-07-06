package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertString(input string) (string, error) {
	trimExp := strings.TrimSpace(input)
	if trimExp == "" {
		return "", errors.New("This input is empty!")
	}
	if isMorse(input) {
		output := morse.ToText(input)
		return output, nil
	}
	output := morse.ToMorse(trimExp)
	return output, nil
}

func isMorse(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return r != '-' && r != ' ' && r != '.'
	})
}
