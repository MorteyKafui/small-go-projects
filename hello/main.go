package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": " שלום עולם ",
	"ur": " ہیلو دنیا ",
	"vi": "Xin chào Thế Giới",
}

func main() {
	var lang string

	flag.StringVar(&lang, "lang", "en", "the required language, eg. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("Unsupported language: %q", l)
	}

	return greeting
}
