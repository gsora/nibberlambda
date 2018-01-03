package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gsora/nibberlambda/nibber"
)

const (
	suh           = "suh my ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
	breathingSuh  = "suh my \xF0\x9F\x85\xB1reathing ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
	clappingNibba = "clapping ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
)

var (
	nibberInstance nibber.Nibber
	apiKey         string
	postURL        string
	numberOfSpaces int = 6
	fillingSpace   string
)

func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
	setupParams()

	if !(numberOfSpaces <= 0 || numberOfSpaces == 1) {
		generateSpaces()
	}

	nibberInstance = nibber.NewNibber(nibber.Emojis)

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(postURL))
	if err != nil {
		log.Fatal(err)
	}

	// We have a strategy here: parse evt into a JSON object, parse it again into a Telegram
	// Update struct, then handle everything like we did before.
	tgObjJSON, err := json.Marshal(evt)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var u tgbotapi.Update
	err = json.Unmarshal(tgObjJSON, &u)
	if err != nil {
		log.Println(err)
		return "", err
	}

	handleUpdate(u, bot)

	return "done!", nil
}

func setupParams() {
	apiKey = os.Getenv("BOT_KEY")
	postURL = os.Getenv("POST_URL")

	if apiKey == "" || postURL == "" {
		log.Fatal("apikey or posturl not found, apikey:", apiKey, "posturl:", postURL)
	}
}

func generateSpaces() {
	fillingSpace = strings.Repeat(" ", numberOfSpaces)
}
