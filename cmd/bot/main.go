package main

import (
	"algoBot/internal/config"
	"algoBot/internal/db"
	"algoBot/internal/db/repository/postgres"
	"algoBot/internal/telegram"
	"algoBot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log/slog"
)

func main() {
	cfg := config.MustConfig()
	log := logger.SetupLogger(cfg.Env)
	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Error("can't create BotApi: " + err.Error())
	}
	botApi.Debug = true

	log.Info("initializing db")
	dataBase := db.InitDB(cfg, log)
	defer dataBase.Close()
	repo := postgres.NewRepository(dataBase)
	bot := telegram.NewBot(botApi, repo, log)

	log.Info("starting app")
	if err := bot.Start(); err != nil {
		log.Error("error while starting app", slog.String("err", err.Error()))
	}
}
