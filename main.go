package main

import (
	"log"
	"time"
	constants "gitlab.ghn.vn/online/integration/bot-telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    var vApiTokenTelegram = "5220166959:AAFRqSJNYwr7pSiD6HpGYsxVqMAktxUpuZU"
	bot, err := tgbotapi.NewBotAPI(vApiTokenTelegram)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
//tetst aaa  aaa
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
        now := time.Now()
        if now.Format("15:04:05")>="17:08"||now.Format("15:04:05")<="08:00"{
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Vui lòng tạo ef giúp team")
            msg.ReplyToMessageID = update.Message.MessageID
            bot.Send(msg)
         }
	}
}