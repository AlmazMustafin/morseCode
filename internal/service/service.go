package service

import (
	"errors"
	"strings"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertAuto(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("нет данных")
	}

	isText := false
	for _, r := range trimmed {
		if (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я') || r == 'ё' || r == 'Ё' {
			isText = true
			break
		}
	}

	if isText {

	return morse.ToMorse(trimmed), nil
	}

	return morse.ToText(trimmed), nil
}
