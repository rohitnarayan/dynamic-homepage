package main

import (
	"encoding/json"
	"io/ioutil"
)

type TranslationMap map[string]map[string]string

// LoadTranslations loads translation data from a JSON file
func LoadTranslations(filePath string) (TranslationMap, error) {
	// Read JSON data from file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into TranslationMap
	translations := make(TranslationMap)
	err = json.Unmarshal(jsonData, &translations)
	if err != nil {
		return nil, err
	}

	return translations, nil
}

// TranslateString translates a string based on the provided translation map and language
func TranslateString(input, language string, translations TranslationMap) string {
	if translation, ok := translations[input]; ok {
		if translatedString, ok := translation[language]; ok {
			return translatedString
		}
	}

	return "Translation not found"
}
