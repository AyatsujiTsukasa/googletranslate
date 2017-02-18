// googletranslate.go

package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

var ()

func main() {
	// Parse input option
	var apiKey string
	//var inputPhrase string
	var targetLanguage string
	flag.StringVar(&apiKey, "apiKey", "", "your api key which you can get form GCP console")
	//flag.StringVar(&inputPhrase, "i", "", "input phrase")
	flag.StringVar(&targetLanguage, "targetLang", "ja", "target language code")
	flag.Parse()

	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		resp, err := client.Translate(ctx, []string{scanner.Text()}, lang, nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp[0].Text)
	}
}
