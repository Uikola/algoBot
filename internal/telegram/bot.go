package telegram

import (
	"algoBot/internal/db/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log/slog"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	repository repository.Repository
	log        *slog.Logger
}

func NewBot(bot *tgbotapi.BotAPI, repository repository.Repository, log *slog.Logger) *Bot {
	return &Bot{bot: bot, repository: repository, log: log}
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
				b.log.Info("handle command error: %s", slog.String("err", err.Error()))
			}
			continue
		}

		if err := b.handleMsg(update.Message); err != nil {
			b.log.Info("handle msg error: %s", slog.String("err", err.Error()))
		}
	}
	return nil
}
