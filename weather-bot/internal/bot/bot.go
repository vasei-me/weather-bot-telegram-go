package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot     *tgbotapi.BotAPI
	handler MessageHandler
}

func NewTelegramBot(token string, handler MessageHandler) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		bot:     bot,
		handler: handler,
	}, nil
}

func (tb *TelegramBot) GetBot() *tgbotapi.BotAPI {
	return tb.bot
}

func (tb *TelegramBot) Start() {
	log.Printf("🤖 ربات با موفقیت راه‌اندازی شد: @%s", tb.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := tb.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Location != nil {
			if err := tb.handler.HandleLocation(update); err != nil {
				log.Printf("Error handling location: %v", err)
			}
			continue
		}

		if err := tb.handler.HandleMessage(update); err != nil {
			log.Printf("Error handling message: %v", err)
		}
	}
}