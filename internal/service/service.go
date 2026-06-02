package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertData(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("переданный файл пуст")
	}

	isMorse := true
	for _, r := range trimmed {
		if r != '.' && r != '-' && r != ' ' && r != '/' && r != '•' && r != '−' {
			isMorse = false
			break
		}
	}

	var result string
	if isMorse {
		result = morse.ToText(trimmed)
	} else {
		result = morse.ToMorse(trimmed)
	}
	return result, nil
}
