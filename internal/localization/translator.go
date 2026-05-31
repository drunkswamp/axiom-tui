package localization

import "fmt"

type Translator struct {
	language Language
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
