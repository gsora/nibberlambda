package main

import (
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gsora/nibberlambda/breath"
	"github.com/gsora/nibberlambda/nibber"
)

func handleUpdate(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.InlineQuery != nil {
		payload := []interface{}{}
		memeStr := nibberInstance.Nibbering(update.InlineQuery.Query)
		clappingMemeStr := strings.Replace(memeStr, " ", " "+nibber.Clap+" ", -1)
		normalMemeString := memeStr
		if numberOfSpaces > 1 {
			normalMemeString = strings.Replace(memeStr, " ", fillingSpace, -1)
		}

		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, suh, normalMemeString)
		article.Description = normalMemeString
		payload = append(payload, article)

		clappingArticle := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID+"-clapping", clappingNibba, clappingMemeStr)
		clappingArticle.Description = clappingMemeStr
		payload = append(payload, clappingArticle)

		breathingMemeStr, err := breath.Breath(memeStr)
		if err != nil {
			log.Printf("[ERROR] cannot breath for request %s\n", update.InlineQuery.Query)
		} else {
			breathingArticle := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID+"-breathing", breathingSuh, breathingMemeStr)
			breathingArticle.Description = breathingMemeStr
			payload = append(payload, breathingArticle)
		}

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       payload,
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	} else if update.Message != nil {
		memeStr := nibberInstance.Nibbering(update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, memeStr)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}

}
