// googletranslate.go

package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "start" {
			break
		}
	}

	fmt.Println("ready to translate")

	// loop until abort
	for scanner.Scan() {
		unquoted, err := strconv.Unquote(scanner.Text())
		if err != nil {
			unquoted = scanner.Text()
		}
		vec, err := base64.StdEncoding.DecodeString(unquoted)
		if err != nil {
			log.Fatal(err)
		}
		var rqst requestFromUnity
		err = json.Unmarshal(vec, &rqst)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Translate(ctx, []string{rqst.OrgText}, lang, nil)
		if err != nil {
			log.Fatal(err)
		}
		var rspnc responceFromUnity
		rspnc.InstanceID = rqst.InstanceID
		rspnc.ResponceText = resp[0].Text
		outputByte, err := json.Marshal(rspnc)
		if err != nil {
			log.Fatal(err)
		}
		result := base64.StdEncoding.EncodeToString(outputByte)
		fmt.Println(result)
	}
}

type requestFromUnity struct {
	InstanceID string
	OrgText    string
}

type responceFromUnity struct {
	InstanceID   string
	ResponceText string
}
