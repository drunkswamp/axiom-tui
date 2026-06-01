package localization

import (
	"fmt"
	"strings"
)

type Translator struct {
	language Language
}

func LanguageFromString(value string) Language {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case string(Russian):
		return Russian
	case string(English):
		return English
	default:
		return English
	}
}

func NewTranslator(language Language) Translator {
	return Translator{language: language}
}

func (t Translator) Language() Language {
	return t.language
}

func (t Translator) T(key Key) string {
	if dictionary, ok := dictionaries[t.language]; ok {
		if value, ok := dictionary[key]; ok {
			return value
		}
	}

	return fmt.Sprintf("[%s]", key)
}
