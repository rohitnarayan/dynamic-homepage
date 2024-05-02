package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Section represents a section of the mobile page with translated strings
type Section struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// MobilePage represents the entire mobile page
type MobilePage struct {
	PageTitle string    `json:"page_title"`
	Sections  []Section `json:"sections"`
}

// LoadPageFromFile dynamically loads the mobile page JSON from a file
func LoadPageFromFile(filePath string, translations TranslationMap, language string) (*MobilePage, error) {
	// Read JSON data from file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into MobilePage struct
	page := new(MobilePage)
	err = json.Unmarshal(jsonData, &page)
	if err != nil {
		return nil, err
	}

	// Translate page title
	page.PageTitle = TranslateString(page.PageTitle, language, translations)

	// Translate section titles and content
	for i := range page.Sections {
		page.Sections[i].Title = TranslateString(page.Sections[i].Title, language, translations)
		page.Sections[i].Content = TranslateString(page.Sections[i].Content, language, translations)
	}

	return page, nil
}

func main() {
	filePath := "./resources/link.json" // Path to the JSON file representing the mobile page

	// Load translation data from file
	translations, err := LoadTranslations("./resources/translations.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	language := "id_ID" // Language to use for translation

	// Load mobile page from file and translate
	page, err := LoadPageFromFile(filePath, translations, language)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	bytes, err := json.MarshalIndent(page, "", "	")
	if err != nil {
		fmt.Println("error occurred = ", err)
	}

	fmt.Println()
	// Print translated mobile page
	fmt.Printf("%+v\n", string(bytes))
}
