package telegram

import (
	"algoBot/internal/db/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	repository repository.Repository
}

func NewBot(bot *tgbotapi.BotAPI, repository repository.Repository) *Bot {
	return &Bot{bot: bot, repository: repository}
}

func (b *Bot) Start() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if err := b.handleCmd(update.Message); err != nil {
				log.Printf("handle command error: %s", err.Error())
			}
			continue
		}

		if err := b.handleMsg(update.Message); err != nil {
			log.Printf("handle msg error: %s", err.Error())
		}
	}
	return nil
}
