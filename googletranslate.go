// googletranslate.go

package main

import (
	"bufio"
	"flag"
	"fmt"
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
	var targetLanguage string
	flag.StringVar(&apiKey, "apiKey", "", "your api key which you can get form GCP console")
	flag.StringVar(&targetLanguage, "targetLang", "ja", "target language code")
	flag.Parse()

	// Setup client
	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ready to translate")

	// loop until abort cmmand
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		resp, err := client.Translate(ctx, []string{scanner.Text()}, lang, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp[0].Text)
	}
}
