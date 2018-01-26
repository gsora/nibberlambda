package main

import (
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"

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
	numberOfSpaces = 6
	fillingSpace   string
)

func init() {
	setupParams()

	if !(numberOfSpaces <= 0 || numberOfSpaces == 1) {
		generateSpaces()
	}
}

func main() {
	lambda.Start(HandleTelegramMessage)
}

// HandleTelegramMessage handles a Telegram message pushed by AWS to this application
func HandleTelegramMessage(tgMessage tgbotapi.Update) {
	nibberInstance = nibber.NewNibber(nibber.Emojis)

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	handleUpdate(tgMessage, bot)
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
